# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 093K

Tested At: 02 Jun 23 01:03 UTC\
Initial Validity Period: 90 day(s)\
Remaining Validity Period: -245 day(s)\
Subject: CN=SHAKEN 093K, OU=SHAKEN, O=Skywave Wireless Inc, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/093K/26bed0c9-863c-4e31-9972-8ca30a9ea44f.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC9TCCApygAwIBAgIQfI40twOP5ZyBCqX4QXCMRjAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjA3MDExNTM5NTdaFw0yMjA5MjkxNTM5NTZaMFMxCzAJBgNVBAYTAlVTMR0wGwYDVQQKExRTa3l3YXZlIFdpcmVsZXNzIEluYzEPMA0GA1UECxMGU0hBS0VOMRQwEgYDVQQDEwtTSEFLRU4gMDkzSzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABCS6iWS0Iobp2F7%2FVyMEW4DCvL555p64gvjYJckebaWkDdNppqfO0E2yZLiqiq4srRVBBUeJPmmC2LAvYAWSnEqjggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIAgDAdBgNVHQ4EFgQUXfAnXpnDgJ9Zhv6JjYELIrfF%2F68wHwYDVR0jBBgwFoAUu5beMRLN05aZhKQ2MGA811KBfScwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYEMDkzSzAKBggqhkjOPQQDAgNHADBEAiABiKeh34jzed0Cn7mt4VZUGfMdAhI8OoY%2BTRyrGEQ2dQIgO5qj6rR5va%2FiCFYeY7kuk9iZIqqHxq4gySqaNfyr2gY%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 02 Jun 23 01:12 UTC