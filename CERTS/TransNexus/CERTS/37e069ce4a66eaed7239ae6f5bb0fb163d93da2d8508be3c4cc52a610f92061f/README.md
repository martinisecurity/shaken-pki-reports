# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 459J

Tested At: 31 Jan 23 21:39 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -93 day(s)\
Subject: CN=SHAKEN 459J, OU=SHAKEN, O=Altaworx, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/d3d6dbfc-2914-49c7-8f47-d0aa5bd5d764/322d7c5da941a6d008ec9cd6a3ef2ee5.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC6jCCApCgAwIBAgIQVYmZ6xyR7ttAo515IG9KGjAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMjMyMDIzMDFaFw0yMjEwMzAyMDIzMDBaMEcxCzAJBgNVBAYTAlVTMREwDwYDVQQKEwhBbHRhd29yeDEPMA0GA1UECxMGU0hBS0VOMRQwEgYDVQQDEwtTSEFLRU4gNDU5SjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABPJGV08FGRs3WB41OCwGMebTlfWmw5IVzfdHHYlg%2FJouszLTL%2FWdLhhuCbXWlJTdAhIziKsib2Gdq4TnJ1a%2BzHKjggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIAgDAdBgNVHQ4EFgQUwVmDrI48O2iyJFsv224AV65XE6wwHwYDVR0jBBgwFoAUu5beMRLN05aZhKQ2MGA811KBfScwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYENDU5SjAKBggqhkjOPQQDAgNIADBFAiA9TEmoXxLDS411GxjEiQ2I%2BowxNdE6y3Ic%2BvYdOJeg5AIhAMXBfsxSTJwB3S3ylc70kElACcgsXO%2FfNhHYhKXDym2f)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |


Generated: 31 Jan 23 21:50 UTC