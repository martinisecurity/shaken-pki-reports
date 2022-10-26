# STIR/SHAKEN CA Ecosystem Compliance

## 1RouteGroup
Name: e_sti_subject_cn\
Source: ATIS-1000080v4\
Citation: ATIS-1000080.v004 / 6.4.1 STI Certificate Requirements\
Effective Date: 16 Jan 22 00:00 UTC\
Description: The Common Name attribute of an End-Entity certificate shall contain the text string “SHAKEN”, followed by a single space, followed by the SPC value identified in the TNAuthList of the End-Entity certificate

### Leaf Certificates

| Status | Subject | Link | Details |
|--------|---------|------|---------|
| error | CN=MYPBXManager SHAKEN, O=MYPBXManager LLC, ST=New York, C=US | [view](../a3872afd09406d2745d204893b6b52bbf6380f84/README.md) | Cannot get SPC value from the TNAuthList extension, bad TNAuthorizationList, bad TNAuthorizationList ASN.1 raw, asn1: syntax error: data truncated |


Generated: 26/10/2022 at 20:32:17