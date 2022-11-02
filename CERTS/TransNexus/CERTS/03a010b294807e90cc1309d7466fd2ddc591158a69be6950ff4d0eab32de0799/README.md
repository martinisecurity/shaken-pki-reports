# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Fusion Connect SHAKEN 2720

Tested At: 02 Nov 22 17:24 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 148 day(s)\
Subject: CN=Fusion Connect SHAKEN 2720, OU=Fusion Connect, O=TransNexus, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA1, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://prod001-cr.rbbnidhub.com/8ZJdVFg7gz/2

[View certificate details](https://understandingwebpki.com/?cert=MIICoTCCAkegAwIBAgIQSFscNJOjGceOfFyuBLk%2BVzAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMTAeFw0yMjAzMzAxMjU0MTVaFw0yMzAzMzAxMjU0MTRaMGAxCzAJBgNVBAYTAlVTMRMwEQYDVQQKEwpUcmFuc05leHVzMRcwFQYDVQQLEw5GdXNpb24gQ29ubmVjdDEjMCEGA1UEAxMaRnVzaW9uIENvbm5lY3QgU0hBS0VOIDI3MjAwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAT63jwuwao3XgseWltjK2ryDbERqdg1YPJq5ERkdas77TTVUuYvquN%2FKmHQd5tGNxd0HFT74V7uRnF%2B0%2Faql%2B59o4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBQgC9%2Fhoul3TB6S5BONMObenIn0jjAfBgNVHSMEGDAWgBSUhjk%2F5PWSoJ%2F%2F3Cd1GppG8HnhYjAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMBYGCCsGAQUFBwEaBAowCKAGFgQyNzIwMAoGCCqGSM49BAMCA0gAMEUCIDNwXtIcpzboL8SV5qUwPjA3OiDBDuLVOeuxTqtMfwpQAiEAt7LcSEYFhKVzHAXNrJ9oJPpw8%2Fzrflp1JEJ6FudrYKU%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |


Generated: 02 Nov 22 17:25 UTC