# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 551G

Tested At: 11 Jan 23 20:38 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -89 day(s)\
Subject: CN=SHAKEN 551G, OU=SHAKEN, O=Brightlink Communications LLC, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/551G/f221d48c-a740-448f-b89c-ad81a6f2c256.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC%2FzCCAqWgAwIBAgIQe8qP3jQagrcIT6W0y4t9yDAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMDcwNzEwMDVaFw0yMjEwMTQwNzEwMDRaMFwxCzAJBgNVBAYTAlVTMSYwJAYDVQQKEx1CcmlnaHRsaW5rIENvbW11bmljYXRpb25zIExMQzEPMA0GA1UECxMGU0hBS0VOMRQwEgYDVQQDEwtTSEFLRU4gNTUxRzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABC%2FJzK4wJ8lNP4zAfNcujY9iXFZkRya7EYoB4hZC0aPFX7YkeBBoljbZUsx0BLJB6vTiRd9qGcleQpUi%2BldivVqjggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIAgDAdBgNVHQ4EFgQUrj8Pkcz8Q6C7uFYrFwVd2%2BvmLrMwHwYDVR0jBBgwFoAUu5beMRLN05aZhKQ2MGA811KBfScwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYENTUxRzAKBggqhkjOPQQDAgNIADBFAiEA8KCdcgEYkJ6wkcFNAFYZCbKglNORY0VatyQdJoPCNzYCIDPWBNE2FoQFAXe%2FRSwCIUWA5ocL5gyJYCl7PDNO9RId)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |


Generated: 11 Jan 23 21:04 UTC