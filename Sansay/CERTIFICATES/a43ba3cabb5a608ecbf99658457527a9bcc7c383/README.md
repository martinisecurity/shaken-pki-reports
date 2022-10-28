# STIR/SHAKEN CA Ecosystem Compliance
## Sansay

### Certificate a43ba3cabb5a608ecbf99658457527a9bcc7c383
Tested At: 2022-10-28 18:22:40 +0000 UTC\
Initial Validity Period: 31 day(s)\
Remaining Validity Period: 15 day(s)\
Subject: CN=SHAKEN  XCast Labs 689J, OU=XCast Labs, O=XCast Labs, ST=California, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US

Link: https://cr.xcastlabs.net/1668234600/xclsshaken.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIDgDCCAyagAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkTwYwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAxMjA2MzAwMFoXDTIyMTExMjA2MzAwMFowbjELMAkGA1UEBhMCVVMxEzARBgNVBAgMCkNhbGlmb3JuaWExEzARBgNVBAoMClhDYXN0IExhYnMxEzARBgNVBAsMClhDYXN0IExhYnMxIDAeBgNVBAMMF1NIQUtFTiAgWENhc3QgTGFicyA2ODlKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAECSaRM8oTiHWzZxnWX07VK9ClMrfygu659qCz3%2FNdIZetgYKmv2%2Fi%2B%2F3bRH4jjvo5INmhtTlMMWWnkQ0T%2Bn9R4qOCAYgwggGEMBYGCCsGAQUFBwEaBAowCKAGFgQ2ODlKMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUQu1fhldX0A%2ByhR0kbkEsERXho%2B4wgcoGA1UdIwSBwjCBv4AUrNOT9UNDzAq%2BRVgXE32SfNzDAUahgZCkgY0wgYoxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRIwEAYDVQQHDAlTYW4gRGllZ28xGzAZBgNVBAoMElNhbnNheSBDb3Jwb3JhdGlvbjESMBAGA1UECwwJU2Fuc2F5IENBMSEwHwYDVQQDDBhTSEFLRU4gU2Fuc2F5IFJvb3QgQ0EgVVOCFBS1XzgF9fB7E7X4sN7tIPJRcD6cMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNIADBFAiAWRwm82JXVfMML9jfOTcpTwuVRfgein8Xj5VtVPC%2BTUwIhAJheo8MaIjn9CnOIuwoUHx3ukXlt9VN0ECJQQMsMxgdW)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_cp1_3_subject_sn | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_sti_subject_cn | error | ATIS-1000080 | Common name shall contain the text string 'SHAKEN 689J' |
| w_pki_subject_rdn_unknown | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| e_cp1_3_ambiguous_identifier | error | United States SHAKEN CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| e_sti_certificate_policies | error | ATIS-1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |

Generated: 28/10/2022 at 18:22:55