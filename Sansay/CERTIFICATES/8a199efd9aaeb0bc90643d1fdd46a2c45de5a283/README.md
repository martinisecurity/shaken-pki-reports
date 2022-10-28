# STIR/SHAKEN CA Ecosystem Compliance
## Sansay

### Certificate 8a199efd9aaeb0bc90643d1fdd46a2c45de5a283
Tested At: 2022-10-28 18:15:29 +0000 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 348 day(s)\
Subject: CN=SHAKEN Current Calls\\, LLC 746J, OU=CurrentCalls, O=Current Calls\\, LLC, ST=California, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US

Link: https://cr.sansay.com/Current_Calls_LLC_746J

View: [Click to view](https://understandingwebpki.com/?cert=MIIC4jCCAomgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkTyEwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAxMTE3MTc0MloXDTIzMTAxMTE3MTc0MlowfzELMAkGA1UEBhMCVVMxEzARBgNVBAgMCkNhbGlmb3JuaWExGzAZBgNVBAoMEkN1cnJlbnQgQ2FsbHMsIExMQzEVMBMGA1UECwwMQ3VycmVudENhbGxzMScwJQYDVQQDDB5TSEFLRU4gQ3VycmVudCBDYWxscywgTExDIDc0NkowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASCnQGHwC%2F7ZcJlA2%2F9NySkl3gZRbr3sReeCrF3k7%2B738HVJR34PYR1YI62ujddP2OrhGjFbg%2BrN8jkJAwgOBfjo4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQ3NDZKMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUS%2FMxhkVL43GV%2BIfIwBzC5TvsU6swHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0cAMEQCIC%2FFmsNMTykUKky%2BQIBMwQP23CknTc%2F4fdiaDLAMV0ZlAiBqkCpc1IS3Sj3ErMoyzhMVT4gVTx97VI1T%2FFCACDE8WQ%3D%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_sti_subject_cn | error | ATIS-1000080 | Common name shall contain the text string 'SHAKEN 746J' |
| w_pki_subject_rdn_unknown | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| e_cp1_3_subject_sn | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_sti_certificate_policies | error | ATIS-1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| e_cp1_3_ambiguous_identifier | error | United States SHAKEN CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |

Generated: 28/10/2022 at 18:15:47