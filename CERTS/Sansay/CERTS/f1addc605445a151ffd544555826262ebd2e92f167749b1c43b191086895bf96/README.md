# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Broadband Dynamics Canada Telecom and Software ULC 884J

Tested At: 15 Dec 22 18:11 UTC\
Initial Validity Period: 162 day(s)\
Remaining Validity Period: 140 day(s)\
Subject: CN=SHAKEN Broadband Dynamics Canada Telecom and Software ULC 884J, OU=Engineering, O=Broadband Dynamics Canada Telecom and Software ULC, ST=Arizona, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA Canada 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://canada-cr.sansay.com/Broadband_Dynamics_Canada_Telecom_and_Software_ULC_884J

[View certificate details](https://understandingwebpki.com/?cert=MIIDFTCCArqgAwIBAgIUVVx8cHEeOCDwuOHerm%2FzJiJkM7UwCgYIKoZIzj0EAwIwgYQxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEvMC0GA1UEAwwmU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgQ2FuYWRhIDEwHhcNMjIxMTIyMjA1NjIzWhcNMjMwNTAzMjA1NjIzWjCBuzELMAkGA1UEBhMCVVMxEDAOBgNVBAgMB0FyaXpvbmExOzA5BgNVBAoMMkJyb2FkYmFuZCBEeW5hbWljcyBDYW5hZGEgVGVsZWNvbSBhbmQgU29mdHdhcmUgVUxDMRQwEgYDVQQLDAtFbmdpbmVlcmluZzFHMEUGA1UEAww%2BU0hBS0VOIEJyb2FkYmFuZCBEeW5hbWljcyBDYW5hZGEgVGVsZWNvbSBhbmQgU29mdHdhcmUgVUxDIDg4NEowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQ3hz0em5WnA6GT8YMyng84jbOKpNqKFTFGpD8BrHkbSQ1oYKlbuZrJRauIatOxf%2BquAMhLdbIOA51%2Bh0GoIPcRo4HQMIHNMBYGCCsGAQUFBwEaBAowCKAGFgQ4ODRKMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUwzBsMFQiskB24Qs%2BigpwZqC4CHAwHwYDVR0jBBgwFoAU4y75ZnRQUjkl%2FyZ8%2BItXynhFsD4wPAYDVR0fBDUwMzAxoC%2BgLYYraHR0cHM6Ly9zdGlwYS1jc3RnYS5jY2lkLm5ldXN0YXIvYXBpL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNJADBGAiEA0CNcLKUa39wOZUdGeAKJj2r9uHxWtMSO64C2kYQP6IkCIQDLhr0jwJ3kbpTyj7ke1RPK%2FnSsCYs5QZRChVpv6Plk9w%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 884J' |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 15 Dec 22 18:35 UTC