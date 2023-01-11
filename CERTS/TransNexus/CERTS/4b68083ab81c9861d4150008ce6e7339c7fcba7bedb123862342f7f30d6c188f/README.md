# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 089K

Tested At: 11 Jan 23 20:37 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -85 day(s)\
Subject: CN=SHAKEN 089K, OU=SHAKEN, O=Logista Solutions, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/5f7c135d-caf3-4661-abb1-d1720e7872e2/11619920d7fee7122cd5c5a1ea6203f4.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC8jCCApmgAwIBAgIQaCBLck5u8A7sW%2B5Og9hRZTAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMTExODM0NTFaFw0yMjEwMTgxODM0NTBaMFAxCzAJBgNVBAYTAlVTMRowGAYDVQQKExFMb2dpc3RhIFNvbHV0aW9uczEPMA0GA1UECxMGU0hBS0VOMRQwEgYDVQQDEwtTSEFLRU4gMDg5SzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABIQ88yUUV71TGLGHW1cNztsi7fYYrbLpReBGpvt%2Boh5vtdDtbcBQDE7Zfsjqtcbb%2BrSpXy77r4bTIGdkkY1q9YCjggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIAgDAdBgNVHQ4EFgQUFKXahMtq7Jv4Ze24WiH3iDiuidowHwYDVR0jBBgwFoAUu5beMRLN05aZhKQ2MGA811KBfScwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYEMDg5SzAKBggqhkjOPQQDAgNHADBEAiAG%2FXdP0bjbkAApnQXtI%2B1%2FFtUOV%2BqotfggC9y04kztoAIgQiQ%2BcVhQFdTY5qQcLwO2x4exIE%2FnLPTGWd0L1Rurt0w%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 11 Jan 23 21:04 UTC