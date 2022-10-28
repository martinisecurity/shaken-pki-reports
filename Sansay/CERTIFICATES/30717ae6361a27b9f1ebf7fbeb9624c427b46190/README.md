# STIR/SHAKEN CA Ecosystem Compliance
## Sansay

### Certificate 30717ae6361a27b9f1ebf7fbeb9624c427b46190
Tested At: 2022-10-28 18:15:17 +0000 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 29 day(s)\
Subject: CN=SHAKEN Broadband Dynamics LLC 583j, OU=Operations, O=Broadband Dynamics LLC, ST=Arizona, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US

Link: https://cr.sansay.com/583j_BROADBAND_DYNAMICS_STIR_SHAKEN.crt

View: [Click to view](https://understandingwebpki.com/?cert=MIIDlTCCAzugAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkU3IwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAyNjIyMjY0MloXDTIyMTEyNTIyMjY0MlowgYIxCzAJBgNVBAYTAlVTMRAwDgYDVQQIDAdBcml6b25hMR8wHQYDVQQKDBZCcm9hZGJhbmQgRHluYW1pY3MgTExDMRMwEQYDVQQLDApPcGVyYXRpb25zMSswKQYDVQQDDCJTSEFLRU4gQnJvYWRiYW5kIER5bmFtaWNzIExMQyA1ODNqMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEGvpK77E6nNmb5Ib3%2B0atBc1rvOUU0j%2FHa5SXcvxc%2F1iNaQpDhxr6Cu8bDw4U2WQdGr4vPdmitNornKsYZ0qtjaOCAYgwggGEMBYGCCsGAQUFBwEaBAowCKAGFgQ1ODNqMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUiYHRBUDlqk8R8LsC1JaQPodW5t8wgcoGA1UdIwSBwjCBv4AUrNOT9UNDzAq%2BRVgXE32SfNzDAUahgZCkgY0wgYoxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRIwEAYDVQQHDAlTYW4gRGllZ28xGzAZBgNVBAoMElNhbnNheSBDb3Jwb3JhdGlvbjESMBAGA1UECwwJU2Fuc2F5IENBMSEwHwYDVQQDDBhTSEFLRU4gU2Fuc2F5IFJvb3QgQ0EgVVOCFBS1XzgF9fB7E7X4sN7tIPJRcD6cMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNIADBFAiEAsUBAzRe4SzfkL4obv3928s8M87Ws9MrvwnycNnp1uJ0CIDYSUZvQ8SdqvaRZiU%2BbH%2BvgnKhv%2BaRlWov%2BaPGa7zyg)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_cp1_3_ambiguous_identifier | error | United States SHAKEN CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| e_sti_certificate_policies | error | ATIS-1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| e_sti_subject_cn | error | ATIS-1000080 | Common name shall contain the text string 'SHAKEN 583j' |
| e_cp1_3_subject_sn | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| w_pki_subject_rdn_unknown | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |

Generated: 28/10/2022 at 18:15:47