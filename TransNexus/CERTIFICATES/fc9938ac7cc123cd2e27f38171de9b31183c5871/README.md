# STIR/SHAKEN CA Ecosystem Compliance
## TransNexus

### Certificate fc9938ac7cc123cd2e27f38171de9b31183c5871
Tested At: 2022-10-28 18:22:54 +0000 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: 1 day(s)\
Subject: CN=SHAKEN 551G, OU=SHAKEN, O=Brightlink Communications LLC, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US

Link: https://certificates.transnexus.com/551G/943338f5-bdb6-4d1e-a761-98bb9d714f94.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIC%2FjCCAqWgAwIBAgIQTyc9ehIrQql%2BHCKh4%2BPttjAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMjIwMDIxNDlaFw0yMjEwMjkwMDIxNDhaMFwxCzAJBgNVBAYTAlVTMSYwJAYDVQQKEx1CcmlnaHRsaW5rIENvbW11bmljYXRpb25zIExMQzEPMA0GA1UECxMGU0hBS0VOMRQwEgYDVQQDEwtTSEFLRU4gNTUxRzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABFidUko4xWrGFhGUEiBZWUjjlFTYe2gxXgB51%2FgEGpt0pdO%2FIxpsWVt6B%2BpTqE%2FU6Awxf8efq4r1GBODyjUqaU6jggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIAgDAdBgNVHQ4EFgQUKCYsEzbzjETRemA0Mt6d9FVq4a4wHwYDVR0jBBgwFoAUu5beMRLN05aZhKQ2MGA811KBfScwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYENTUxRzAKBggqhkjOPQQDAgNHADBEAiBIjy7sxTT6NOIT%2BpGwKvVV5sXDbAB92GuIUyFGNjUh5AIgMfdD5gHgcXNzaTzlbteIV5SLZsYTFi83K6iz4iGsEA0%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| w_pki_subject_rdn_unknown | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| e_incorrect_ku_encoding | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| e_cp1_3_subject_sn | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |

Generated: 28/10/2022 at 18:22:55