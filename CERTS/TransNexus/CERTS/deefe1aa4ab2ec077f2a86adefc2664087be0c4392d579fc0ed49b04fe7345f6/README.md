# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 291K

Tested At: 30 Jan 23 23:00 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -88 day(s)\
Subject: CN=SHAKEN 291K, OU=SHAKEN, O=Hypercore Networks\\, Inc, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/b66a7496-cdd4-46e3-b219-cb0ecdc80d22/dd9ef96fbf98f0935e2baf5a0813aabc.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC%2BjCCAp%2BgAwIBAgIQTwjHxlOjwl%2BvnA8%2Bwy3lojAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMjcxNjMyMzlaFw0yMjExMDMxNjMyMzhaMFYxCzAJBgNVBAYTAlVTMSAwHgYDVQQKExdIeXBlcmNvcmUgTmV0d29ya3MsIEluYzEPMA0GA1UECxMGU0hBS0VOMRQwEgYDVQQDEwtTSEFLRU4gMjkxSzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABMAoKzNqBbe%2BvnyQBVeA1HpBhaJaoHswFTdgkD8kb%2F0ZGFk8FRBE61SycN3UISHqWEbF6%2BmC7Upkv0P%2B3hQT3tejggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAdBgNVHQ4EFgQUe9vW1IHL0rHm2TFPiufhDkE9f4MwHwYDVR0jBBgwFoAUu5beMRLN05aZhKQ2MGA811KBfScwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYEMjkxSzAKBggqhkjOPQQDAgNJADBGAiEA%2BCLtWuLIP3A%2BTd%2BqO4zZyzngtoY24RS1P4VqN%2F%2BWXn0CIQCHl4Su6AbxD391ISa7Krgpqsx1LcfeY6tokWcdCTLKhg%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 30 Jan 23 23:10 UTC