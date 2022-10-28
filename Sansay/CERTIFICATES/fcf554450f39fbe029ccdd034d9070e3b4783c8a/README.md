# STIR/SHAKEN CA Ecosystem Compliance
## Sansay

### Certificate fcf554450f39fbe029ccdd034d9070e3b4783c8a
Tested At: 2022-10-28 19:22:09 +0000 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 362 day(s)\
Subject: CN=SHAKEN Ytel Inc. 703J, OU=Carrier, O=Ytel Inc., ST=California, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US

Link: https://cr.sansay.com/Ytel_Inc._703J

View: [Click to view](https://understandingwebpki.com/?cert=MIIDejCCAyCgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkUuEwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAyNDIxMTE0MFoXDTIzMTAyNDIxMTE0MFowaDELMAkGA1UEBhMCVVMxEzARBgNVBAgMCkNhbGlmb3JuaWExEjAQBgNVBAoMCVl0ZWwgSW5jLjEQMA4GA1UECwwHQ2FycmllcjEeMBwGA1UEAwwVU0hBS0VOIFl0ZWwgSW5jLiA3MDNKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEziv%2FlNgdVlsf6Zexob0C2KIXHxGqlRm%2Fi5%2F7gKFlF%2Fzz4nEpTjYYwFBAU729P0bDYEYRQX7n4zUxsHsRnM77CqOCAYgwggGEMBYGCCsGAQUFBwEaBAowCKAGFgQ3MDNKMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQU1E0NNUJWulI%2BlrfRkApPVOln0IowgcoGA1UdIwSBwjCBv4AUrNOT9UNDzAq%2BRVgXE32SfNzDAUahgZCkgY0wgYoxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRIwEAYDVQQHDAlTYW4gRGllZ28xGzAZBgNVBAoMElNhbnNheSBDb3Jwb3JhdGlvbjESMBAGA1UECwwJU2Fuc2F5IENBMSEwHwYDVQQDDBhTSEFLRU4gU2Fuc2F5IFJvb3QgQ0EgVVOCFBS1XzgF9fB7E7X4sN7tIPJRcD6cMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNIADBFAiArjT3%2BPyqqyU%2BQu54vm%2BsXZRzjd4DVsMTzeTUwzoL%2FeQIhAN%2F4k98nsL%2BUzkXaNOQV9ClH1COiLcRJMjgZ%2B05Z53zA)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_cp1_3_subject_sn | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_cp1_3_ambiguous_identifier | error | United States SHAKEN CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| e_sti_subject_cn | error | ATIS-1000080 | Common name shall contain the text string 'SHAKEN 703J' |
| w_pki_subject_rdn_unknown | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| e_sti_certificate_policies | error | ATIS-1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |

Generated: 28/10/2022 at 19:22:10