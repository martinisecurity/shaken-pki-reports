# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 918J

Tested At: 16 Mar 23 19:03 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -174 day(s)\
Subject: CN=SHAKEN 918J, OU=SHAKEN, O=VoIPX International Inc, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/3bbe1f4c-3184-44ee-84f4-2f63891aa57b/d9b86d264d8bfdcea8367f36c90b1ee6.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC%2BTCCAp%2BgAwIBAgIQXWUz7ha2JdOPcOU9IqtIsTAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjA5MTUxOTQxMjhaFw0yMjA5MjIxOTQxMjdaMFYxCzAJBgNVBAYTAlVTMSAwHgYDVQQKExdWb0lQWCBJbnRlcm5hdGlvbmFsIEluYzEPMA0GA1UECxMGU0hBS0VOMRQwEgYDVQQDEwtTSEFLRU4gOTE4SjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABGKx%2B8FLgWyoVcVXSHpCaq%2FN6Tn3KmAFp4cziuo9AQN5A8WR%2BigRk37bxwfaBqnEb1%2BVmCqf%2B%2F74G8U15LQz6mKjggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIAgDAdBgNVHQ4EFgQUaKT4l67xfF3j83%2FGZRYlz3qX2jUwHwYDVR0jBBgwFoAUu5beMRLN05aZhKQ2MGA811KBfScwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYEOTE4SjAKBggqhkjOPQQDAgNIADBFAiB3CG2xtoF61duhLwO9p46XejMbES84QYe7cMrh9UclUwIhAKKfySQ2Rpgw3%2FZmNmP%2B3ddgXapagFxmPCjaercsqj9y)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |


Generated: 16 Mar 23 19:18 UTC