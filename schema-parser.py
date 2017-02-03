#!/usr/bin/env python

from __future__ import unicode_literals, print_function, division

import argparse
import json
import re
import sys

try:
    import requests
except ImportError:
    requests = None

debug = False


def dprint(*args):
    if debug:
        dprint(*args)

go_types = {"string": "string",
            "date-time": "string",
            "integer": "int",
            "boolean": "bool",
            "null": "nil",
            "enum": "string",
            # "entity": "map[string]interface{}",
            "entity": "entity",
            "object": "map[string]interface{}",
            "array": "[]interface{}"}

GO_ENTITY_RE = r"\s+(\w+)\s+\bentity\b.*json"

# Needs "." to include newlines
PYTHON_ENDPOINT_RE = r"class (\w+).*?super.*?args\)"


def parse_schema(schema):
    version = None
    results = []
    for endpoint, ebody in schema.items():

        epdict = {"name": endpoint}
        dprint()
        dprint("Endpoint: ", endpoint)

        if type(ebody) == dict:
            epdict["ops"] = []
            for op, body in ebody.items():
                opdict = {"op": op}
                dprint("Operation: ", op)

                opdict["entity"] = body.get("entity")

                if ("bodyParamSchema" in body and
                        "properties" in body["bodyParamSchema"]):
                    opdict["params"] = []
                    for param, pbody in body[
                            "bodyParamSchema"]["properties"].items():
                        pdict = {"name": param}
                        dprint("Param:", param)
                        try:
                            tp = pbody["type"]
                            if type(tp) == list:
                                tp = "string"
                            pdict["type"] = tp
                            dprint("Type: ", tp)
                        except KeyError:
                            if "enum" in pbody:
                                pdict["type"] = "enum"
                                pdict["enum"] = pbody["enum"]
                                dprint("Type: Enum")
                            else:
                                raise
                        opdict["params"].append(pdict)
                if ("returnParamSchema" in body and "properties"
                        in body["returnParamSchema"]):
                    opdict["return_params"] = []
                    for param, pbody in body[
                            "returnParamSchema"]["properties"].items():
                        pdict = {"name": param}
                        dprint("Param:", param)
                        try:
                            tp = pbody["type"]
                            if type(tp) == list:
                                tp = "string"
                            pdict["type"] = tp
                            dprint("Type: ", tp)
                        except KeyError:
                            if "enum" in pbody:
                                pdict["type"] = "enum"
                                pdict["enum"] = pbody["enum"]
                                dprint("Type: Enum")
                            elif "containment" in pbody:
                                pdict["type"] = "entity"
                                dprint("Type: Entity")
                            else:
                                raise
                        opdict["return_params"].append(pdict)
                epdict["ops"].append(opdict)

        elif endpoint == "version":
            version = ebody
        results.append(epdict)

    for result in results:
        result["version"] = version

    return results


def snake_to_cap_camel(s):
    return "".join((elem.title() for elem in s.split("_")))


def python_name_from_path(path):
    result = path.replace(":", "")

    parts = filter(lambda x: x, result.split("/"))

    for index, part in enumerate(parts):
        if part in ["name", "id"]:
            parts[index] = parts[index - 1].rstrip("s")
    result = snake_to_cap_camel("_".join(parts)).replace("Name", "")
    return result


def make_python(pschema):

    top = """#!/usr/bin/env python
\"\"\"
Provides the various v2.1 endpoints and entities
\"\"\"
from .base import Endpoint as _Endpoint
# from .base import ContainerEndpoint as _ContainerEndpoint
# from .base import SingletonEndpoint as _SingletonEndpoint
# from .base import SimpleReferenceEndpoint as _SimpleReferenceEndpoint
# from .base import ListEndpoint as _ListEndpoint
# from .base import StringEndpoint as _StringEndpoint
# from .base import MetricCollectionEndpoint as _MetricCollectionEndpoint
# from .base import MetricEndpoint as _MetricEndpoint
from .base import Entity as _Entity
# from .base import MonitoringStatusEndpoint as _MonitoringStatusEndpoint
__copyright__ = "Copyright 2017, Datera, Inc."

"""
    ep_kls = ("\n\nclass {}Endpoint(_Endpoint):\n"
              "    _path_spec = \"{}\"\n\n"
              "    def __init__(self, *args):\n"
              "        super({}Endpoint, self).__init__(*args)")

    en_kls = ("\n\nclass {}Entity(_Entity):\n"
              "    _path_spec = \"{}\"\n\n"
              "    def __init__(self, *args):\n"
              "        super({}Entity, self).__init__(*args)")

    sep = "        self._set_subendpoint({}Endpoint)"
    snt = "        self._entity_cls({}Entity)"
    # sne = "        self._entity_cls({}Endpoint)"

    made = set()
    result = {}
    result = {}
    eparts_dict = {}

    for endpoint in pschema:
        ###########################################################
        # TODO(_alastor_): remove this fix when the tenant endpoint
        # Spec is fixed
        if "((" in endpoint["name"]:
            endpoint["name"] = endpoint["name"].split("((")[0]
        ###########################################################
        endpoint["pname"] = endpoint["name"].split("/")[
            -1]
        eparts_dict[endpoint["name"]] = endpoint["name"].split("/")

    # Entity Creation
    en_pschema = filter(lambda x: ":" in x["pname"], pschema)
    for entity in en_pschema:
        name = python_name_from_path(entity["name"])
        if name not in made:
            result[entity["name"]] = en_kls.format(
                name,
                endpoint["name"],
                name)
        made.add(name)

    # Endpoint Creation
    ep_pschema = filter(lambda x: ":" not in x["pname"], pschema)
    for endpoint in ep_pschema:
        name = python_name_from_path(endpoint["name"])
        if name not in made:
            spec = eparts_dict[endpoint["name"]]
            result[endpoint["name"]] = ep_kls.format(
                name,
                endpoint["name"],
                name)
        made.add(name)

    # Subendpoints/Entity processing
    for spec, epstring in result.items():
        # Entities can have endpoints, but don't act as endpoints
        eparts = eparts_dict[spec]
        above = None
        if len(eparts) >= 2:
            above = "/".join(eparts[:-1])
        topline = filter(lambda x: x, epstring.splitlines())[0].lower()
        if above and "entity" not in topline:
            try:
                result[above] = "\n".join((
                    result[above],
                    sep.format(python_name_from_path(spec))))
            except KeyError:
                pass
        if above and "entity" in topline:
            try:
                result[above] = "\n".join((
                    result[above],
                    snt.format(python_name_from_path(spec))))
            except KeyError:
                pass
        # elif above and ":" not in spec:
        #     try:
        #         result[above] = "\n".join((
        #             result[above],
        #             sne.format(python_name_from_path(spec))))
        #     except KeyError:
        #         pass
    # Returning concatenated top section + the classes sorted by their names
    return "".join((top, "\n".join(
        [elem[1] for elem in sorted(result.items(), key=lambda x: x[0])])))


def make_golang(pschema, package):

    # Create initial set of structs
    # After they're created we need to iterate through and link up related
    # Entities
    made = set()
    struct_registry = set()
    result = "package {}\n".format(package)

    struct = "type %s struct {\n%s\n}"

    for endpoint in pschema:
        if "ops" not in endpoint:
            continue
        if "[" in endpoint["name"]:
            endpoint["name"] = endpoint["name"].split("[")[0]
        end = endpoint["name"].split("/")[-1]
        name = end.replace(":", "").replace("(", "").replace(
            ")", "").replace("?", "")
        name = snake_to_cap_camel(name)
        if name in made:
            continue
        made.add(name)
        for opdict in endpoint["ops"]:
            subname = "".join((opdict["op"].title(), name))
            params = ""
            for param in opdict.get("params", {}):
                params = "\n".join((
                    params,
                    "\t{}\t{}\t`json:\"{},omitempty\"`".format(
                        snake_to_cap_camel(param["name"]),
                        go_types[param["type"]],
                        param["name"])))
            result = "\n".join((result, "// " + endpoint["name"]))
            result = "\n".join((result, struct % (subname, params)))
            struct_registry.add(subname)

            rparams = ""
            for param in opdict.get("return_params", {}):
                rparams = "\n".join((
                    rparams,
                    "\t{}\t{}\t`json:\"{},omitempty\"`".format(
                        snake_to_cap_camel(param["name"]),
                        go_types[param["type"]],
                        param["name"])))
            subname = "".join(("Return", subname))
            result = "\n".join((result, "// " + endpoint["name"]))
            result = "\n".join((result, struct % (subname, rparams)))
            struct_registry.add(subname)

    # Now we'll link up related entities

    resultlines = result.splitlines()
    for index, line in enumerate(resultlines):
        match = re.match(GO_ENTITY_RE, line)
        if match:
            struct = match.group(1)
            name = "ReturnRead" + struct
            if name in struct_registry:
                resultlines[index] = line.replace("entity", name)
            else:
                resultlines[index] = line.replace(
                    "entity", "map[string]interface{}")
    return "\n".join(resultlines)


def make_golang_v2(pschema, package):
    pass


def _go_ep():
    pass


def _go_en():
    pass


def _go_enep():
    pass


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

    result = None

    # import ipdb; ipdb.set_trace()
    if args.language == "golang":
        result = make_golang(pschema, args.package)
    elif args.language == "python":
        result = make_python(pschema)

    if args.output:
        with open(args.output, 'w+') as f:
            f.write(result)
    else:
        print(result)

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
