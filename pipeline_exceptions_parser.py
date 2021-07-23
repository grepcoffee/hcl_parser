import hcl
import json

def open_parse_exep(entivoment):
    with open('filename', 'r') as exceptions_file:
        exceptions = json.load(exceptions_file)
    for exception in exceptions:
        if exception['entivoment'] == entivoment:
            exception_details = exception['exception_details'][0]
    return exception_details

def policy_matching(policy_name, environment):
    exp_details = open_parse_exep(environment)
    matched_policies = []
    for exception in exp_details.keys():
        for policies in policy_name.keys():
            if exception in policies:
                matched_policies.append(policies)
    return matched_policies

####################################
############### MAIN ###############
####################################

policy = 'sentinel.hcl'

with open(policy, 'r+') as sentinelhcl_file:
    exceptions = open_parse_exep(e)
    sentinelhcl = hcl.load(sentinelhcl_file)
    hcl_policies = sentinelhcl['policy']
    matched_policies = policy_matching(hcl_policies, e)
    for i in hcl_policies:
        if i in matched_policies:
            enforcement_level = exceptions[i]
            hcl_policies[i]['enforcement_level'] = enforcement_level
    hcl_parsed = hcl.dumps(hcl_policies)
    sentinelhcl_file.write(hcl_parsed)
