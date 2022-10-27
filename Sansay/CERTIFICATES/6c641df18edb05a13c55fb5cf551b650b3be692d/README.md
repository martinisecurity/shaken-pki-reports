# STIR/SHAKEN CA Ecosystem Compliance
## Sansay

### Certificate 6c641df18edb05a13c55fb5cf551b650b3be692d
Tested At: 2022-10-27 21:26:02 +0000 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 21 day(s)\
Subject: emailAddress=jhansen@ztelco.com, CN=SHAKEN Zray Technologies Corporation 862J, OU=Ztelco, O=Zray Technologies Corporation, ST=California, C=US, emailAddress=jhansen@ztelco.com\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US

Link: https://cr.sansay.com/862J/order/261_862J_58

View: [Click to view](https://understandingwebpki.com/?cert=MIIDxjCCA2ugAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkUaYwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAxODIwMTc1NFoXDTIyMTExNzIwMTc1NFowgbIxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMSYwJAYDVQQKDB1acmF5IFRlY2hub2xvZ2llcyBDb3Jwb3JhdGlvbjEPMA0GA1UECwwGWnRlbGNvMTIwMAYDVQQDDClTSEFLRU4gWnJheSBUZWNobm9sb2dpZXMgQ29ycG9yYXRpb24gODYySjEhMB8GCSqGSIb3DQEJARYSamhhbnNlbkB6dGVsY28uY29tMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEGwlDbaswwZmUeqGiVZFghXZ6vA0esmcuWAkunny4Y5JeOWy4o7zGeMh8F0bmzHnM6pDRMwcziGauzxH7wuyYq6OCAYgwggGEMBYGCCsGAQUFBwEaBAowCKAGFgQ4NjJKMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUL0b0tJbf48H7Hpij0jc1Te5soLkwgcoGA1UdIwSBwjCBv4AUrNOT9UNDzAq%2BRVgXE32SfNzDAUahgZCkgY0wgYoxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRIwEAYDVQQHDAlTYW4gRGllZ28xGzAZBgNVBAoMElNhbnNheSBDb3Jwb3JhdGlvbjESMBAGA1UECwwJU2Fuc2F5IENBMSEwHwYDVQQDDBhTSEFLRU4gU2Fuc2F5IFJvb3QgQ0EgVVOCFBS1XzgF9fB7E7X4sN7tIPJRcD6cMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNJADBGAiEArGgiRrWhHIiK0Hkx3RBSLCuWIZ3JbdnQQ%2B%2Bbm14uESoCIQDfOJreBQpUzTPEiaKPjj9ifM5PIFqZPqHZN5rJU0G1AA%3D%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_cp1_3_subject_sn | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_sti_certificate_policies | error | ATIS-1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| e_sti_subject_cn | error | ATIS-1000080 | Common name shall contain the text string 'SHAKEN 862J' |
| w_cp_1_3_subject_email | warn | United States SHAKEN CP | Email addresses are not allowed as the CP does not specify how to validate them |
| w_cp1_3_subject_rdn_unknown | warn | United States SHAKEN CP | Only CN, C, O, and SERIALNUMBER can be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| e_cp1_3_ambiguous_identifier | error | United States SHAKEN CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |

Generated: 27/10/2022 at 21:27:34