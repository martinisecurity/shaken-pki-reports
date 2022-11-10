# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Threshold Communications Inc 563J

Tested At: 10 Nov 22 06:41 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 16 day(s)\
Subject: CN=SHAKEN Threshold Communications Inc 563J, OU=NOC, O=Threshold Communications Inc, ST=Washington, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/Threshold_Communications_Inc_563J

[View certificate details](https://understandingwebpki.com/?cert=MIIC7zCCApWgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkU3cwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAyNjIzMjcxOVoXDTIyMTEyNTIzMjcxOVowgYoxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApXYXNoaW5ndG9uMSUwIwYDVQQKDBxUaHJlc2hvbGQgQ29tbXVuaWNhdGlvbnMgSW5jMQwwCgYDVQQLDANOT0MxMTAvBgNVBAMMKFNIQUtFTiBUaHJlc2hvbGQgQ29tbXVuaWNhdGlvbnMgSW5jIDU2M0owWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARMkuoAaHJ0IZKIi3iifdRfyX26s7V3xg0p9qopk0q0CYe1hyb1jEicNMVhXBk%2BWX0GcRj50k81FqyGwPoEvOXSo4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQ1NjNKMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUCtMVhpspSBvylbdCt200kXSDv7UwHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0gAMEUCIB2eiplC3AfNzqkk%2BiAuMN2xxe8v5%2B0RXec%2BG79GKQ9uAiEAonDmJ40N%2BiYe9IAO5XgI0ldOfdBoQaoYh9aLki6iPvE%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 563J' |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 10 Nov 22 06:43 UTC