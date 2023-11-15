# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Current Calls, LLC 746J

Tested At: 15 Nov 23 18:03 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 364 day(s)\
Subject: CN=SHAKEN Current Calls\\, LLC 746J, OU=CurrentCalls, O=Current Calls\\, LLC, ST=California, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/Current_Calls_LLC_746J

[View certificate details](https://understandingwebpki.com/?cert=MIIC4zCCAomgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkhjwwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMTExNDEzMzQxMVoXDTI0MTExMzEzMzQxMVowfzELMAkGA1UEBhMCVVMxEzARBgNVBAgMCkNhbGlmb3JuaWExGzAZBgNVBAoMEkN1cnJlbnQgQ2FsbHMsIExMQzEVMBMGA1UECwwMQ3VycmVudENhbGxzMScwJQYDVQQDDB5TSEFLRU4gQ3VycmVudCBDYWxscywgTExDIDc0NkowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASCnQGHwC%2F7ZcJlA2%2F9NySkl3gZRbr3sReeCrF3k7%2B738HVJR34PYR1YI62ujddP2OrhGjFbg%2BrN8jkJAwgOBfjo4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQ3NDZKMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUS%2FMxhkVL43GV%2BIfIwBzC5TvsU6swHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0gAMEUCIQCMow6xfSKwVGnDHwHnLrBSUci2A2NDSbKZ%2BeAQeMlfcgIgAKmnQBp0fTN8om6vrfcONkogXEOW%2FHCG3hzxYgConks%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 746J' |


Generated: 15 Nov 23 18:10 UTC