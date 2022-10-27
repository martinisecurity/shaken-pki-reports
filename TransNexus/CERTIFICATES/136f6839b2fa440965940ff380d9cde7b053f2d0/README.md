# STIR/SHAKEN CA Ecosystem Compliance
## TransNexus

### Certificate 136f6839b2fa440965940ff380d9cde7b053f2d0
Tested At: 2022-10-27 18:24:41 +0000 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 154 day(s)\
Subject: CN=Fusion Connect SHAKEN 2720, OU=Fusion Connect, O=TransNexus, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA1, OU=SHAKEN, O=TransNexus\\, Inc., C=US

Link: https://prod001-cr.rbbnidhub.com/8ZJdVFg7gz/2

View: [Click to view](https://understandingwebpki.com/?cert=MIICoTCCAkegAwIBAgIQSFscNJOjGceOfFyuBLk%2BVzAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMTAeFw0yMjAzMzAxMjU0MTVaFw0yMzAzMzAxMjU0MTRaMGAxCzAJBgNVBAYTAlVTMRMwEQYDVQQKEwpUcmFuc05leHVzMRcwFQYDVQQLEw5GdXNpb24gQ29ubmVjdDEjMCEGA1UEAxMaRnVzaW9uIENvbm5lY3QgU0hBS0VOIDI3MjAwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAT63jwuwao3XgseWltjK2ryDbERqdg1YPJq5ERkdas77TTVUuYvquN%2FKmHQd5tGNxd0HFT74V7uRnF%2B0%2Faql%2B59o4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBQgC9%2Fhoul3TB6S5BONMObenIn0jjAfBgNVHSMEGDAWgBSUhjk%2F5PWSoJ%2F%2F3Cd1GppG8HnhYjAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMBYGCCsGAQUFBwEaBAowCKAGFgQyNzIwMAoGCCqGSM49BAMCA0gAMEUCIDNwXtIcpzboL8SV5qUwPjA3OiDBDuLVOeuxTqtMfwpQAiEAt7LcSEYFhKVzHAXNrJ9oJPpw8%2Fzrflp1JEJ6FudrYKU%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_cp1_3_subject_sn | error | CPv1.3 | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| w_cp1_3_subject_rdn_unknown | warn | CPv1.3 | Only CN, C, O, and SERIALNUMBER can be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| e_sti_crl_distribution | error | ATIS-1000080v4 | CRL Distribution Point shall be reachable if the requesting IP address within the program ACLs |
| e_incorrect_ku_encoding | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| e_sti_certificate_policies | error | ATIS-1000080v4 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| e_cp1_3_ambiguous_identifier | error | CPv1.3 | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |

Generated: 27/10/2022 at 18:24:52