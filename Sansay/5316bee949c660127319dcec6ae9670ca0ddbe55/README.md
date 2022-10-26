# STIR/SHAKEN CA Ecosystem Compliance
## Sansay

### Certificate 5316bee949c660127319dcec6ae9670ca0ddbe55
Tested At: 2022-10-26 20:21:29 +0000 UTC\
Initial Validity Period: 31 day(s)\
Remaining Validity Period: 10 day(s)\
Subject: CN=SHAKEN Consolidated Communications 5113, OU=NOC, O=Consolidated Communications, ST=New Hampshire, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US

Link: https://cr.sansay.com/Consolidated_5113.crt

View: [Click to view](https://understandingwebpki.com/?cert=MIIDnjCCA0SgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkTNwwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAwNTE1MjExN1oXDTIyMTEwNTE1MjExN1owgYsxCzAJBgNVBAYTAlVTMRYwFAYDVQQIDA1OZXcgSGFtcHNoaXJlMSQwIgYDVQQKDBtDb25zb2xpZGF0ZWQgQ29tbXVuaWNhdGlvbnMxDDAKBgNVBAsMA05PQzEwMC4GA1UEAwwnU0hBS0VOIENvbnNvbGlkYXRlZCBDb21tdW5pY2F0aW9ucyA1MTEzMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEIk3QPdu6bR%2BCslRuPfG5Xa3JWRiGOg1SH71pp9phvaR0Q%2Bot4r%2B62VqKUStE5bzD80dVHyhtP7o3X5%2F7tI7rTaOCAYgwggGEMBYGCCsGAQUFBwEaBAowCKAGFgQ1MTEzMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUWWvtJS74qwN3g%2FEc0ZmIHpxzySwwgcoGA1UdIwSBwjCBv4AUrNOT9UNDzAq%2BRVgXE32SfNzDAUahgZCkgY0wgYoxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRIwEAYDVQQHDAlTYW4gRGllZ28xGzAZBgNVBAoMElNhbnNheSBDb3Jwb3JhdGlvbjESMBAGA1UECwwJU2Fuc2F5IENBMSEwHwYDVQQDDBhTSEFLRU4gU2Fuc2F5IFJvb3QgQ0EgVVOCFBS1XzgF9fB7E7X4sN7tIPJRcD6cMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNIADBFAiEAiyP4jLqJ9truCjxo7DQorIgRcuxa%2FricArGlRf3I%2BsUCIEoJjjypP3FPCEv1ASs6DMAheDLE8lSKYA0ck6XRvNBs)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_cp1_3_ambiguous_identifier | error | CPv1.3 | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| e_cp1_3_subject_sn | error | CPv1.3 | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_sti_subject_cn | error | ATIS-1000080v4 | Common name shall contain the text string 'SHAKEN 5113' |
| e_sti_certificate_policies | error | ATIS-1000080v4 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| w_cp1_3_subject_rdn_unknown | warn | CPv1.3 | Only CN, C, O, and SERIALNUMBER can be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| e_sti_crl_distribution | error | ATIS-1000080v4 | CRL Distribution Point shall be reachable if the requesting IP address within the program ACLs |

Generated: 26/10/2022 at 20:22:11