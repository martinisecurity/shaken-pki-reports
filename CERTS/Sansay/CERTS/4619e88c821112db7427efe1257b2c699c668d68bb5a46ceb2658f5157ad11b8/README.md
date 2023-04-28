# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Cloud Connect LLC 455K

Tested At: 28 Apr 23 02:09 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 362 day(s)\
Subject: CN=SHAKEN Cloud Connect LLC 455K, OU=1234, O=Cloud Connect LLC, ST=New York, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/Cloud_Connect_LLC_455K

[View certificate details](https://understandingwebpki.com/?cert=MIIC1zCCAn2gAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkbJQwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDQyNTAwMzQxNVoXDTI0MDQyNDAwMzQxNVowczELMAkGA1UEBhMCVVMxETAPBgNVBAgMCE5ldyBZb3JrMRowGAYDVQQKDBFDbG91ZCBDb25uZWN0IExMQzENMAsGA1UECwwEMTIzNDEmMCQGA1UEAwwdU0hBS0VOIENsb3VkIENvbm5lY3QgTExDIDQ1NUswWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATJC%2Bmlwk8xBqPsmtIQiBfoW4CtEWBEEe2EZg359g6IK4HmqT05LXACh6xp0k8gqZoEs9FFZ0eGX4eB9Ip8Jpato4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQ0NTVLMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUf%2FElel0H5gpfmW6NpziiHUjs6Q0wHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0gAMEUCIHzcXhaCpVI%2BxdEm4xpmCBxfnaVmSf49M1bbOFJERmH%2BAiEA65QL%2FyVNqW8Y2svU7BylsMLwu%2BSt1dyhKcV8SgKgbbM%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 455K' |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 28 Apr 23 02:17 UTC