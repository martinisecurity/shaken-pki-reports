# STIR/SHAKEN CA Ecosystem Compliance

## Certificate DISH Wireless L.L.C.SHAKEN.490J

Tested At: 21 Aug 23 20:17 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 87 day(s)\
Subject: L=Denver, ST=CO, OU=Security Operations, O=Dish Wireless, C=US, CN=DISH Wireless L.L.C.SHAKEN.490J\
Issuer: L=Lowell, ST=Massachusettes, OU=US, O=NetNumber Inc, C=US, CN=NetNumber SHAKEN Root Intermediate CA 1\
Link: https://stsh.dish-wireless.com/pubsti.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIDEDCCApagAwIBAgIIc9ARFiikZY0wCgYIKoZIzj0EAwIwgY4xMDAuBgNVBAMMJ05ldE51bWJlciBTSEFLRU4gUm9vdCBJbnRlcm1lZGlhdGUgQ0EgMTELMAkGA1UEBhMCVVMxFjAUBgNVBAoMDU5ldE51bWJlciBJbmMxCzAJBgNVBAsMAlVTMRcwFQYDVQQIDA5NYXNzYWNodXNldHRlczEPMA0GA1UEBwwGTG93ZWxsMB4XDTIyMTExNjE3MDc1NVoXDTIzMTExNjE3MDc1NVowgYsxKDAmBgNVBAMMH0RJU0ggV2lyZWxlc3MgTC5MLkMuU0hBS0VOLjQ5MEoxCzAJBgNVBAYTAlVTMRYwFAYDVQQKDA1EaXNoIFdpcmVsZXNzMRwwGgYDVQQLDBNTZWN1cml0eSBPcGVyYXRpb25zMQswCQYDVQQIDAJDTzEPMA0GA1UEBwwGRGVudmVyMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEojGDE7ZSf7g3Kun4SSg61eXsyBaHqrd9oTtTpGvDPx0Un6Nn1mU17bTYLbpE9OYCpLB2TkOYQT9gMcmsJYamSqOB3jCB2zAWBggrBgEFBQcBGgQKMAigBhYENDkwSjAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAfBgNVHSMEGDAWgBRxL8iC3KjgIuPfoGj5%2BF5chN7lvTAdBgNVHQ4EFgQUvRoWnvBvj3k0jBlCBGO59XkrhX8wRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMBoGA1UdIAEB%2FwQQMA4wDAYKYIZIAYb%2FCQEBATAKBggqhkjOPQQDAgNoADBlAjEA%2FPQYlRsBfU6eay%2B5R%2BSH%2B4zpwnbsZe1UzmUdhPir7vLEjvr3O8%2F9E1V2CRGX1BeaAjBrQE46byngB%2B0T7FZkKH9zyU%2FFO2%2BN0UnqgjuFN9s6qjZco%2FpK854fKD3gS12JExg%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [n_atis_certificate_policy_critical](../../ISSUES/n_atis_certificate_policy_critical/README.md) | notice | ATIS1000080 | STI certificates should contain a CertificatePolicies extension marked uncritical |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 490J' |


Generated: 21 Aug 23 20:18 UTC