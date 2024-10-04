# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Edify SHAKEN

Tested At: 04 Oct 24 16:25 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: -163 day(s)\
Subject: CN=Edify SHAKEN, O=Edify, ST=California, C=US\
Issuer: CN=1RouteGroup SHAKEN Intermediate CA, O=1RouteGroup, ST=Texas, C=US\
Link: https://sti.edifylabs.net/04-24-2023.crt

[View certificate details](https://x509.io/?cert=MIIChDCCAiqgAwIBAgICEPMwCgYIKoZIzj0EAwIwYDELMAkGA1UEBhMCVVMxDjAMBgNVBAgMBVRleGFzMRQwEgYDVQQKDAsxUm91dGVHcm91cDErMCkGA1UEAwwiMVJvdXRlR3JvdXAgU0hBS0VOIEludGVybWVkaWF0ZSBDQTAeFw0yMzA0MjQyMjUzNDBaFw0yNDA0MjMyMjUzNDBaMEkxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMQ4wDAYDVQQKDAVFZGlmeTEVMBMGA1UEAwwMRWRpZnkgU0hBS0VOMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAErD%2BQzjYKUa1SS5jpvgpmi4rpfDLO%2BqL00GCaHz5xmIw5xaY%2FJqx%2FWDVbw08z%2F%2B8q6VuuxOQ6YpbH2ON0h5KAhKOB6jCB5zAVBggrBgEFBQcBGgQJMAigBhY4NjdKMAwGA1UdEwEB%2FwQCMAAwHQYDVR0OBBYEFEfOgfVkej1N9sBDovIeUnuBtfSuMIGQBgNVHSMEgYgwgYWAFEUJJWxdLsfhUJ4ghhO5JMZzT%2BI%2BoWmkZzBlMQswCQYDVQQGEwJVUzEOMAwGA1UECAwFVGV4YXMxETAPBgNVBAcMCEZ0IFdvcnRoMREwDwYDVQQKDAhHQlNEVGVjaDEgMB4GA1UEAwwXR0JTRFRlY2ggU0hBS0VOIFJvb3QgQ0GCAhAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNIADBFAiEAsC4LgbyDVMZO2giW315KzKfIoqTGCKT3nO7dsdepd7sCIG9OryXkI0kzp2UOLCu9eDY%2FBNb3kU5a0Jy6g8anYsIN)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_tn_auth_list_spc_format](../../ISSUES/e_atis_tn_auth_list_spc_format/README.md) | error | ATIS1000080 | bad TNAuthorizationList, bad TNAuthorizationList ASN.1 raw, asn1: syntax error: data truncated |
| [e_atis_tn_auth_list](../../ISSUES/e_atis_tn_auth_list/README.md) | error | ATIS1000080 | bad TNAuthorizationList, bad TNAuthorizationList ASN.1 raw, asn1: syntax error: data truncated |
| [e_atis_serial_number_size](../../ISSUES/e_atis_serial_number_size/README.md) | error | ATIS1000080 | STI certificates shall have a serial number that contains at least 64 bits, got 13 |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Cannot get SPC value from the TNAuthList extension, bad TNAuthorizationList, bad TNAuthorizationList ASN.1 raw, asn1: syntax error: data truncated |
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension does not contain a single OID value that identifies the SHAKEN Certificate Policy established by the STI-PA |
| [e_atis_ext_crl_distribution](../../ISSUES/e_atis_ext_crl_distribution/README.md) | error | ATIS1000080 | STI End-Entity certificates shall contain a CRL Distribution Points extension |
| [e_atis_ext_certificate_policies](../../ISSUES/e_atis_ext_certificate_policies/README.md) | error | ATIS1000080 | The Certificate Policies extension is not present |


Generated: 04 Oct 24 16:29 UTC