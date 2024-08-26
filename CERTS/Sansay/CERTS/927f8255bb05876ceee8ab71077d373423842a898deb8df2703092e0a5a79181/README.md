# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Airespring 996H

Tested At: 26 Aug 24 18:36 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -93 day(s)\
Subject: CN=SHAKEN Airespring 996H, OU=Airespring NOC, O=Airespring, ST=California, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/996H/429C7C70711E3820F0B8E1DEAE6FF32622649A2C.pem

[View certificate details](https://x509.io/?cert=MIIC1DCCAnugAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkmiwwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTI0MDQyNTEzMjE1MloXDTI0MDUyNTEzMjE1MlowcTELMAkGA1UEBhMCVVMxEzARBgNVBAgMCkNhbGlmb3JuaWExEzARBgNVBAoMCkFpcmVzcHJpbmcxFzAVBgNVBAsMDkFpcmVzcHJpbmcgTk9DMR8wHQYDVQQDDBZTSEFLRU4gQWlyZXNwcmluZyA5OTZIMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEO%2FDtGCxkRKOrL%2FCO4UTvqQyw1p4MqYMIqk92awEPU6Zj%2BIbwKfhahHQHw%2BYo3o%2BOQupLUMCFOnweTLm%2B2rEXzKOB2zCB2DAWBggrBgEFBQcBGgQKMAigBhYEOTk2SDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFGXP4f0G2020cYTsBTAPzTqgd8AUMB8GA1UdIwQYMBaAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNHADBEAiBLRO24aiP9mO3JqYuzW31YHJEqbTaD9cY89mT9HTkqfAIgKKj8TkPABD0jQ8mpPuC9JOpm8zJ83%2BC9ZGA2eeqESpg%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 996H', but common name is 'SHAKEN Airespring 996H' |


Generated: 26 Aug 24 18:49 UTC