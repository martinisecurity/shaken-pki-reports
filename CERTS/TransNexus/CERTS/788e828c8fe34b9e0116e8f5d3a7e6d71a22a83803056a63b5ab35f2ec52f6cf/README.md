# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 012K

Tested At: 21 Aug 23 20:01 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -334 day(s)\
Subject: CN=SHAKEN 012K, OU=SHAKEN, O=CallCurrent\\, Inc., C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/06ebd24a-1f0a-46d5-8a2f-a7ae49be8ed6/08d04d1bfffa8ddf52073599e552f0bc.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC8jCCApmgAwIBAgIQZddlodPFcS38uh03AYtMKTAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjA5MTQwODM2MjdaFw0yMjA5MjEwODM2MjZaMFAxCzAJBgNVBAYTAlVTMRowGAYDVQQKExFDYWxsQ3VycmVudCwgSW5jLjEPMA0GA1UECxMGU0hBS0VOMRQwEgYDVQQDEwtTSEFLRU4gMDEySzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABGyMU2qSbbBMCrFUh7Du5lKkv8v32jQhkjfHfLfypWX6sMuXU65k557rX0jBuGUZqFTqvqLyTI43L6CumDAwK46jggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIAgDAdBgNVHQ4EFgQUZt2FAwu8%2BfUQz7lyrlvOgbmd0nowHwYDVR0jBBgwFoAUu5beMRLN05aZhKQ2MGA811KBfScwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYEMDEySzAKBggqhkjOPQQDAgNHADBEAiB4Q4Fu3o%2FjM44%2F2Qio2krOGDdKq%2FKULB4elfN8AKtaiAIgH8cTEhsFfCg%2FZcrdKHO2o5QYWZmTvSzqq52wEKkV4tk%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 21 Aug 23 20:18 UTC