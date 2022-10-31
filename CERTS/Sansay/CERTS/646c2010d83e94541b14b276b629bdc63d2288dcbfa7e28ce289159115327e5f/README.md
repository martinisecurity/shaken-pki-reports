# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Zray Technologies Corporation 862J

Tested At: 31 Oct 22 20:25 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 12 day(s)\
Subject: emailAddress=jhansen@ztelco.com, CN=SHAKEN Zray Technologies Corporation 862J, OU=Ztelco, O=Zray Technologies Corporation, ST=California, C=US, emailAddress=jhansen@ztelco.com\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/862J/order/255_862J_58

View: [Click to view](https://understandingwebpki.com/?cert=MIIDxDCCA2ugAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkT7IwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAxMjIwNDc1NFoXDTIyMTExMTIwNDc1NFowgbIxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMSYwJAYDVQQKDB1acmF5IFRlY2hub2xvZ2llcyBDb3Jwb3JhdGlvbjEPMA0GA1UECwwGWnRlbGNvMTIwMAYDVQQDDClTSEFLRU4gWnJheSBUZWNobm9sb2dpZXMgQ29ycG9yYXRpb24gODYySjEhMB8GCSqGSIb3DQEJARYSamhhbnNlbkB6dGVsY28uY29tMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEGwlDbaswwZmUeqGiVZFghXZ6vA0esmcuWAkunny4Y5JeOWy4o7zGeMh8F0bmzHnM6pDRMwcziGauzxH7wuyYq6OCAYgwggGEMBYGCCsGAQUFBwEaBAowCKAGFgQ4NjJKMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUL0b0tJbf48H7Hpij0jc1Te5soLkwgcoGA1UdIwSBwjCBv4AUrNOT9UNDzAq%2BRVgXE32SfNzDAUahgZCkgY0wgYoxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRIwEAYDVQQHDAlTYW4gRGllZ28xGzAZBgNVBAoMElNhbnNheSBDb3Jwb3JhdGlvbjESMBAGA1UECwwJU2Fuc2F5IENBMSEwHwYDVQQDDBhTSEFLRU4gU2Fuc2F5IFJvb3QgQ0EgVVOCFBS1XzgF9fB7E7X4sN7tIPJRcD6cMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNHADBEAiAar%2FuqiwU3flXH%2BubdimIIHRyXbc15dHqy1gx7Nb%2Fm4gIgS5QXbHIAn394FIlQGbMwHUdlWZX8DTiKe7E0X1MmUlA%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_cp1_3_ambiguous_identifier](../../ISSUES/e_cp1_3_ambiguous_identifier/README.md) | error | United States SHAKEN CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [w_cp_1_3_subject_email](../../ISSUES/w_cp_1_3_subject_email/README.md) | warn | United States SHAKEN CP | Email addresses are not allowed as the CP does not specify how to validate them |
| [e_cp1_3_subject_sn](../../ISSUES/e_cp1_3_subject_sn/README.md) | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_sti_certificate_policies](../../ISSUES/e_sti_certificate_policies/README.md) | error | ATIS&#x2011;1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [w_pki_subject_rdn_unknown](../../ISSUES/w_pki_subject_rdn_unknown/README.md) | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_sti_subject_cn](../../ISSUES/e_sti_subject_cn/README.md) | error | ATIS&#x2011;1000080 | Common name shall contain the text string 'SHAKEN 862J' |


Generated: 31/10/2022 at 20:32:42