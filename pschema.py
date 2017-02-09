#!/usr/bin/env python

from __future__ import unicode_literals, print_function, division

import argparse
import abc
import json
# import re
import sys

try:
    import requests
except ImportError:
    requests = None

try:
    import inflection
except ImportError:
    inflection = None

debug = False

go_types = {"string": "string",
            "date-time": "string",
            "integer": "int",
            "boolean": "bool",
            "null": "nil",
            "enum": "string",
            # "entity": "map[string]interface{}",
            # "entity": "entity",
            "object": "map[string]string",
            "array": "[]interface{}"}


def snake_to_cap_camel(s):
    return "".join((elem.title() for elem in s.split("_")))


def singularize(s):
    tmp = s.lower()
    result = None
    if tmp not in ["metadata", "dns"]:
        result = inflection.singularize(s)
    return result or s


class ApiWriter(object):

    def __init__(self, separate_files=False):
        self.separate_files = separate_files

    @abc.abstractmethod
    def module_header(self):
        pass

    @abc.abstractmethod
    def entity_header(self):
        pass

    @abc.abstractmethod
    def endpoint_header(self):
        pass

    @abc.abstractmethod
    def entity_endpoint_header(self):
        pass

    @abc.abstractmethod
    def entity(self, name, attrs):
        pass

    @abc.abstractmethod
    def endpoint(self, endpoint_data):
        pass

    @abc.abstractmethod
    def entity_endpoint(self, entity_endpoint_data):
        pass


class GoApiWriter(ApiWriter):

    en_func_template = """
func (en {en_name}Entity) Reload() ({en_name}Entity, error) {{
\tvar n {en_name}Entity
\tr, _ := conn.Get(en.Path)
\td, e, err := getData(r)
\tif e.Message != "" {{
\t\treturn n, errors.New(strings.Join(append([]string{{e.Message}}, e.Errors...), ":"))
\t}}
\tif err != nil {{
\t\tpanic(err)
\t}}
\terr = json.Unmarshal(d, &n)
\treturn n, nil
}}
func (en {en_name}Entity) Set(bodyp ...string) ({en_name}Entity, error) {{
\tvar n {en_name}Entity
\tr, _ := conn.Put(en.Path, false, bodyp...)
\td, e, err := getData(r)
\tif e.Message != "" {{
\t\treturn n, errors.New(strings.Join(append([]string{{e.Message}}, e.Errors...), ":"))
\t}}
\tif err != nil {{
\t\tpanic(err)
\t}}
\terr = json.Unmarshal(d, &n)
\treturn n, nil
}}
func (en {en_name}Entity) Delete(bodyp ...string) error {{
\tr, _ := conn.Delete(en.Path, bodyp...)
\t_, e, err := getData(r)
\tif e.Message != "" {{
\t\treturn errors.New(strings.Join(append([]string{{e.Message}}, e.Errors...), ":"))
\t}}
\tif err != nil {{
\t\tpanic(err)
\t}}
\treturn nil
}}
"""
    ep_new_template = """
func New{ep_name}Endpoint(parent string) {ep_name}Endpoint {{
\tpath := strings.Trim(strings.Join([]string{{parent, "{path}"}}, "/"), "/")
\treturn {ep_name}Endpoint{{
\t\tPath: path,
\t\t{attrs}
\t}}
}}
"""
    ep_func_template = """
func (ep {ep_name}Endpoint) Create(bodyp ...string) ({en_name}Entity, error) {{
\tvar en {en_name}Entity
\tr, _ := conn.Post(ep.Path, bodyp...)
\td, e, err := getData(r)
\tif e.Message != "" {{
\t\treturn en, errors.New(strings.Join(append([]string{{e.Message}}, e.Errors...), ":"))
\t}}
\tif err != nil {{
\t\tpanic(err)
\t}}
\terr = json.Unmarshal(d, &en)
\tif err != nil {{
\t\tpanic(err)
\t}}
\treturn en, nil
}}
func (ep {ep_name}Endpoint) List(queryp ...string) ([]{en_name}Entity, error) {{
\tvar ens []{en_name}Entity
\tr, _ := conn.Get(ep.Path, queryp...)
\td, e, err := getData(r)
\tif e.Message != "" {{
\t\treturn ens, errors.New(strings.Join(append([]string{{e.Message}}, e.Errors...), ":"))
\t}}
\tif err != nil {{
\t\tpanic(err)
\t}}
\terr = json.Unmarshal(d, &ens)
\tif err != nil {{
\t\tpanic(err)
\t}}
\treturn ens, nil
}}
"""

    def __init__(self):
        super(GoApiWriter, self).__init__(False)

    def module_header(self):
        return (
            """package dsdk

import (
    "encoding/json"
    //"fmt"
    "strings"
    "errors"
)

// Using a global here because screw having to pass this around to everything
// even via an autogeneration script.  We may hit some limitations later with
// concurrency, but I need this working now.
var conn *ApiConnection

type RootEp struct {
\tPath string
\tAppInstances AppInstancesEndpoint
\tApi ApiEndpoint
\tAppTemplates AppTemplatesEndpoint
\tInitiators InitiatorsEndpoint
\tInitiatorGroups InitiatorGroupsEndpoint
\tAccessNetworkIpPools AccessNetworkIpPoolsEndpoint
\tStorageNodes StorageNodesEndpoint
\tSystem SystemEndpoint
\tEventLogs EventLogsEndpoint
\tAuditLogs AuditLogsEndpoint
\tFaultLogs FaultLogsEndpoint
\tRoles RolesEndpoint
\tUsers UsersEndpoint
\tUpgrade UpgradeEndpoint
\tTime TimeEndpoint
\tTenants TenantsEndpoint
}

func NewRootEp(hostname, port, username, password, apiVersion, tenant, timeout string, headers map[string]string, secure bool) (*RootEp, error) {
\tvar err error
\t//Initialize global connection object
\tconn, err = NewApiConnection(hostname, port, username, password, apiVersion, tenant, timeout, headers, secure)
\tif err != nil {
\t\treturn nil, err
\t}
\terr = conn.Login()
\tif err != nil {
\t\treturn nil, err
\t}
\treturn &RootEp{
\t\tPath:         "",
\tAppInstances: NewAppInstancesEndpoint(""),
\tApi: NewApiEndpoint(""),
\tAppTemplates: NewAppTemplatesEndpoint(""),
\tInitiators: NewInitiatorsEndpoint(""),
\tInitiatorGroups: NewInitiatorGroupsEndpoint(""),
\tAccessNetworkIpPools: NewAccessNetworkIpPoolsEndpoint(""),
\tStorageNodes: NewStorageNodesEndpoint(""),
\tSystem: NewSystemEndpoint(""),
\tEventLogs: NewEventLogsEndpoint(""),
\tAuditLogs: NewAuditLogsEndpoint(""),
\tFaultLogs: NewFaultLogsEndpoint(""),
\tRoles: NewRolesEndpoint(""),
\tUsers: NewUsersEndpoint(""),
\tUpgrade: NewUpgradeEndpoint(""),
\tTime: NewTimeEndpoint(""),
\tTenants: NewTenantsEndpoint(""),
\t}, nil
}
""")

    def entity_header(self):
        pass

    def endpoint_header(self):
        pass

    def entity_endpoint_header(self):
        pass

    def entity(self, name, attrs, entities):
        json_template = """`json:"{},omitempty"`"""
        attr_template = "\t{name} {type} {json}"
        attr_list = []
        for attr in attrs:
            attr_name = snake_to_cap_camel(attr['name'])
            attr_type = None
            # Check if type exists as an entity
            if not attr_type and attr['name'] in entities:
                attr_type = snake_to_cap_camel(attr['name']) + "Entity"
            # Check if singularized type exists as an entity
            if (not attr_type and
                    singularize(attr['name']) in entities):
                attr_type = "".join((
                    "[]",
                    singularize(snake_to_cap_camel(attr['name'])),
                    "Entity"))
            if not attr_type:
                attr_type = go_types.get(attr['type'], "interface{}")
            attr_json = json_template.format(attr['name'])
            attrib = attr_template.format(
                name=attr_name, type=attr_type, json=attr_json)
            attr_list.append(attrib)
        result = ("type {name}Entity struct {{\n"
                  "{attrs}\n}}\n").format(
                 name=snake_to_cap_camel(name),
                 attrs="\n".join(sorted(attr_list)))
        result = "\n".join((result,
                            self.en_func_template.format(
                                en_name=snake_to_cap_camel(
                                    singularize(name)))))
        return result

    def endpoint(self, name, attrs, entities):
        attr_list = []
        attr_template = "\t{name} {type}Endpoint"
        new_attr_list = []
        new_attr_template = "\t{name}: New{type}Endpoint(path),"
        for subendpoint in attrs['subep']:
            subendpoint_name = subendpoint.replace(name, "")
            attr_list.append(attr_template.format(
                name=snake_to_cap_camel(subendpoint_name),
                type=snake_to_cap_camel(subendpoint)))
            new_attr_list.append(new_attr_template.format(
                name=snake_to_cap_camel(subendpoint_name),
                type=snake_to_cap_camel(subendpoint)))
        result = ("type {name}Endpoint struct {{\n"
                  "\tPath string\n"
                  "{attrs}\n}}\n").format(
                  name=snake_to_cap_camel(name),
                  attrs="\n".join(sorted(attr_list)))
        # if attrs['type'] == "endpoint":
        en_name = attrs['path'].split("/")[-1]
        if singularize(en_name) in entities:
            result = "\n".join((
                result,
                self.ep_func_template.format(
                    en_name=snake_to_cap_camel(
                        singularize(en_name)),
                    ep_name=snake_to_cap_camel(name))))
        result = "\n".join((
            result,
            self.ep_new_template.format(
                ep_name=snake_to_cap_camel(name),
                path=attrs['path'].strip("/"),
                attrs="\n".join(sorted(new_attr_list)))))
        return result

    def entity_endpoint(self, entity_endpoint_data):
        pass


def parse_schema(schema):
    """
    Returns (endpoint_dict, entity_dict)
    """

    # Filter stream and live operations since they're not useful for
    # sdk operations
    OP_FILTERS = ["stream", "live"]

    results = []

    for endpoint, ep_body in schema.items():
        # TODO(_alastor_): When these weird endpoints are fixed, remove this
        # workaround
        if "(" in endpoint:
            endpoint = endpoint.split("(")[0]
        ep_dict = {"name": endpoint.strip("/").replace(
                           "/", "_").replace(":", ""),
                   "ops": [],
                   "subep": [],
                   "path": endpoint}
        if type(ep_body) == dict:
            for op, op_body in ep_body.items():
                if op in OP_FILTERS:
                    continue
                op_dict = {"op": op,
                           "entity": op_body.get("entity"),
                           "args": [],
                           "return": []}

                body_params = op_body.get("bodyParamSchema", {})
                for param, param_body in body_params.get(
                        "properties", {}).items():
                    param_dict = {"name": param}
                    param_type = param_body.get("type")
                    if type(param_type) == list or "enum" in param_body:
                        param_type = "string"
                    param_dict["type"] = param_type
                    op_dict["args"].append(param_dict)

                return_params = op_body.get("returnParamSchema", {})
                for param, param_body in return_params.get(
                        "properties", {}).items():
                    param_dict = {"name": param}
                    param_type = param_body.get("type")
                    if type(param_type) == list or "enum" in param_body:
                        param_type = "string"
                    elif "containment" in param_body:
                        param_type = "entity"
                    param_dict["type"] = param_type
                    op_dict["return"].append(param_dict)

                ep_dict["ops"].append(op_dict)

            results.append(ep_dict)

    for result in results:
        path = result["path"]
        parts = path.strip("/").split("/")
        parent = "/".join(parts[:-1])
        end = parts[-1]
        if ":" in end:
            result["type"] = "entity"
        for other in results:
            other_path = other["path"]
            other_end = other_path.strip("/").split("/")[-1]
            if parent == other_path.strip("/"):
                other["subep"].append(result["name"])
                if ":" in other_end and not result.get("type"):
                    result["type"] = "endpoint"
                else:  # not result.get("type"):
                    result["type"] = "entity_endpoint"
    for result in results:
        if not result.get("type"):
            result["type"] = "endpoint"

    entity_dict = {}
    for endpoint in results:
        for op in endpoint["ops"]:
            if op["entity"] not in entity_dict and op["entity"]:
                entity_dict[op["entity"]] = op["return"]
            elif not op["entity"]:
                endpoint["type"] = "entity_endpoint"

    return results, entity_dict


def main(args):
    t = None
    if args.url:
        if not requests:
            print("Requests package required when providing API URL")
            sys.exit(1)
        response = requests.get(args.url)
        t = response.text
    elif args.file:
        with open(args.file) as f:
            t = f.read()
    else:
        t = sys.stdin.read()

    pschema = parse_schema(json.loads(t))

    api = GoApiWriter()
    result_str = ""
    result_str = "\n".join((result_str, api.module_header()))
    endpoints, entities = pschema
    entity_strs = []
    for name, attrs in entities.items():
        entity_strs.append(api.entity(name, attrs, entities))
    entity_strs = sorted(entity_strs)
    result_str = "\n".join(((result_str, "\n".join(entity_strs))))

    endpoint_strs = []
    for endpoint in endpoints:
        endpoint_strs.append(api.endpoint(endpoint['name'], endpoint, entities))
    result_str = "\n".join(((result_str, "\n".join(endpoint_strs))))
    print(result_str)


if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    parser.add_argument("-f", "--file", action="store_true",
                        help="File containing schema")
    parser.add_argument("-l", "--language", action="store", default="golang",
                        help="Language to output schema in")
    parser.add_argument("-o", "--output", action="store", default=None,
                        help="Output file")
    parser.add_argument("-u", "--url", action="store", default=None,
                        help="URL where the API schema can be obtained")
    parser.add_argument("-p", "--package", action="store", default="dapi",
                        help="Golang package types should go in")
    args = parser.parse_args()

    sys.exit(main(args))
