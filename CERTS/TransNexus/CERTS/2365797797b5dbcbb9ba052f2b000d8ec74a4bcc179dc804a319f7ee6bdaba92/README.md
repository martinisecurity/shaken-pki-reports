# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 551G

Tested At: 01 Nov 22 20:26 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: 6 day(s)\
Subject: CN=SHAKEN 551G, OU=SHAKEN, O=Brightlink Communications LLC, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/551G/8240af5f-7059-41ee-b8bb-0e42bb96f67c.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIC%2FzCCAqWgAwIBAgIQYQxK5gfO4kgktIK4k4hdyzAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMzExOTQyMDNaFw0yMjExMDcxOTQyMDJaMFwxCzAJBgNVBAYTAlVTMSYwJAYDVQQKEx1CcmlnaHRsaW5rIENvbW11bmljYXRpb25zIExMQzEPMA0GA1UECxMGU0hBS0VOMRQwEgYDVQQDEwtTSEFLRU4gNTUxRzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABBB8ivjGQerUCYiwMAHi6u9qw3tohu%2F3%2BWXjzpmom61J18D%2F9ZaI8UBonCw4ZafJsnZ%2BPzCxz3WjYwjEv%2B3u146jggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAdBgNVHQ4EFgQUA%2F0uFBzhrqzEgL2%2F%2BrT3F5VpfDMwHwYDVR0jBBgwFoAUu5beMRLN05aZhKQ2MGA811KBfScwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYENTUxRzAKBggqhkjOPQQDAgNIADBFAiBAmcwV0WQRjJeIGgbI3bGl4pYrqE6wt1%2Bx9fzh9SjoRgIhAKFlHiPjW5zaBtCRyvXpB0iW0tp1kSZM1U6R0YoxweZb)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 01/11/2022 at 20:34:21