# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 6529

Tested At: 28 Nov 22 20:41 UTC\
Initial Validity Period: 366 day(s)\
Remaining Validity Period: 338 day(s)\
Subject: CN=SHAKEN 6529, OU=Core Network Engineering, O=T-Mobile USA\\, Inc., L=Bellevue, ST=Washington, C=US\
Issuer: CN=TMOBILE-PROD-SUB-STIRSHAKEN-EC, O=TMOBILE-USA, C=US\
Link: https://t-mobile-sticr.fosrvt.com/d1ed2fcf74511801e3df6deffbf762b764221487a8305a7da32e9efab2cbf358.pem

[View certificate details](https://understandingwebpki.com/?cert=MIICsjCCAlegAwIBAgIPAYQzjU8hOw3ZcypWftHaMAoGCCqGSM49BAMCMEwxCzAJBgNVBAYTAlVTMRQwEgYDVQQKEwtUTU9CSUxFLVVTQTEnMCUGA1UEAxMeVE1PQklMRS1QUk9ELVNVQi1TVElSU0hBS0VOLUVDMB4XDTIyMTEwMTEzNDY0OVoXDTIzMTEwMTE0MTY0OVowgYsxCzAJBgNVBAYTAlVTMRMwEQYDVQQIEwpXYXNoaW5ndG9uMREwDwYDVQQHEwhCZWxsZXZ1ZTEbMBkGA1UEChMSVC1Nb2JpbGUgVVNBLCBJbmMuMSEwHwYDVQQLExhDb3JlIE5ldHdvcmsgRW5naW5lZXJpbmcxFDASBgNVBAMTC1NIQUtFTiA2NTI5MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE%2F3hHIIWT1sUmI2nJrHe9D7FdxhRsxMZD6maM5f9gyssKc%2FisbkBkOBGQDRUpi%2FAB42V4AzqFmrprzf%2FdRq1p56OB2zCB2DAfBgNVHSMEGDAWgBTaeENnEVHteHQmMY%2BpWhzvx%2FepWjAdBgNVHQ4EFgQUrDssKVPXXTPM5vFOK8Q%2F5mZE3vAwDgYDVR0PAQH%2FBAQDAgeAMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwFgYIKwYBBQUHARoECjAIoAYWBDY1MjkwDAYDVR0TAQH%2FBAIwADAKBggqhkjOPQQDAgNJADBGAiEAwNJbWzbS%2B2bEyWy2XcdUx7gXn38%2BZ%2B%2Fg9cruvLg4xhgCIQC5Mlv3DCOXoPQwihkBc5D5%2BiOvw0CUJShm8QIqsUwhuQ%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 28 Nov 22 20:41 UTC