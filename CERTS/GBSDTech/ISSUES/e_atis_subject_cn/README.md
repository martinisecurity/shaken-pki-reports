# STIR/SHAKEN CA Ecosystem Compliance

## GBSDTech

Name: e_atis_subject_cn\
Source: ATIS1000080\
Citation: ATIS-1000080.v004 / 6.4.1 STI Certificate Requirements\
Effective Date: 16 Jan 22 00:00 UTC\
Description: The Common Name attribute of an End-Entity certificate shall contain the text string “SHAKEN”, followed by a single space, followed by the SPC value identified in the TNAuthList of the End-Entity certificate

### Leaf Certificates

| Status | Subject | Link | Details |
|--------|---------|------|---------|
| error | FracTEL LLC SHAKEN | [view](../../CERTS/81188428a77465a01dbc3a49509aa0486510c109472f8af6e5cb1887a682040e/README.md) | Cannot get SPC value from the TNAuthList extension, bad TNAuthorizationList, bad TNAuthorizationList ASN.1 raw, asn1: syntax error: data truncated |
| error | Edify SHAKEN | [view](../../CERTS/224a236b36499273c1a11a37e7df69b3dd72f4747dad0f3faf092b0069eb0b1d/README.md) | Cannot get SPC value from the TNAuthList extension, bad TNAuthorizationList, bad TNAuthorizationList ASN.1 raw, asn1: syntax error: data truncated |
| error | MYPBXManager SHAKEN | [view](../../CERTS/0c7bf2cc1741b8036003876afadd109dfd5a6b0fb7af3662ae4d02e3340ad9ce/README.md) | Cannot get SPC value from the TNAuthList extension, bad TNAuthorizationList, bad TNAuthorizationList ASN.1 raw, asn1: syntax error: data truncated |

### CA Certificates

No error, warning, or notice level issues were found


Generated: 17 Dec 22 12:22 UTC