# STIR/SHAKEN CA Ecosystem Compliance

## Certificate FracTEL LLC SHAKEN

Tested At: 31 Jan 23 21:42 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 127 day(s)\
Subject: CN=FracTEL LLC SHAKEN, O=FracTEL LLC, ST=Florida, C=US\
Issuer: CN=1RouteGroup SHAKEN Intermediate CA, O=1RouteGroup, ST=Texas, C=US\
Link: https://ssc.fractel.net/ssc/fractelssc.pem

[View certificate details](https://understandingwebpki.com/?cert=MIICjTCCAjOgAwIBAgICEOYwCgYIKoZIzj0EAwIwYDELMAkGA1UEBhMCVVMxDjAMBgNVBAgMBVRleGFzMRQwEgYDVQQKDAsxUm91dGVHcm91cDErMCkGA1UEAwwiMVJvdXRlR3JvdXAgU0hBS0VOIEludGVybWVkaWF0ZSBDQTAeFw0yMjA2MDcyMTQ0MTlaFw0yMzA2MDcyMTQ0MTlaMFIxCzAJBgNVBAYTAlVTMRAwDgYDVQQIDAdGbG9yaWRhMRQwEgYDVQQKDAtGcmFjVEVMIExMQzEbMBkGA1UEAwwSRnJhY1RFTCBMTEMgU0hBS0VOMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEtw4%2BCDli7CPXfrQ6HogsF7XFgxMcBRAed8Ho5w30VfqD%2B28Pbrsj4Q7fQ8VDRt0cQC3TQ7LAaxjyb%2BfRWYzIyaOB6jCB5zAVBggrBgEFBQcBGgQJMAigBhY5NjVIMAwGA1UdEwEB%2FwQCMAAwHQYDVR0OBBYEFGmtH0KE1euEpn%2FjAld6EP%2Fpr%2FDfMIGQBgNVHSMEgYgwgYWAFEUJJWxdLsfhUJ4ghhO5JMZzT%2BI%2BoWmkZzBlMQswCQYDVQQGEwJVUzEOMAwGA1UECAwFVGV4YXMxETAPBgNVBAcMCEZ0IFdvcnRoMREwDwYDVQQKDAhHQlNEVGVjaDEgMB4GA1UEAwwXR0JTRFRlY2ggU0hBS0VOIFJvb3QgQ0GCAhAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNIADBFAiACuub%2Boy0hrFF3puTRu8JgHLUJned770D%2B%2BfJCi110BQIhAJqou%2Bi8acSeNe9FbQnfVblHXPJdgdRzQJ7zC3z2%2FD6x)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_tn_auth_list](../../ISSUES/e_atis_tn_auth_list/README.md) | error | ATIS1000080 | bad TNAuthorizationList, bad TNAuthorizationList ASN.1 raw, asn1: syntax error: data truncated |
| [e_atis_crl_distribution](../../ISSUES/e_atis_crl_distribution/README.md) | error | ATIS1000080 | STI End-Entity certificates shall contain a CRL Distribution Points extension |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_serial_number](../../ISSUES/e_atis_serial_number/README.md) | error | ATIS1000080 | STI certificates shall include a Serial Number field containing an integer greater than zero. The serial number shall contain at least 64 bits of output from a Cryptographically Secure PseudoRandom Number Generator (CSPRNG) |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Cannot get SPC value from the TNAuthList extension, bad TNAuthorizationList, bad TNAuthorizationList ASN.1 raw, asn1: syntax error: data truncated |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | bad TNAuthorizationList, bad TNAuthorizationList ASN.1 raw, asn1: syntax error: data truncated |


Generated: 31 Jan 23 21:50 UTC