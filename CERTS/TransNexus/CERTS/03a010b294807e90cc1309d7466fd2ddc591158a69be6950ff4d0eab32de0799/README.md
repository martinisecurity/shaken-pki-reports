# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Fusion Connect SHAKEN 2720

Tested At: 21 Nov 23 17:46 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: -236 day(s)\
Subject: CN=Fusion Connect SHAKEN 2720, OU=Fusion Connect, O=TransNexus, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA1, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://prod001-cr.rbbnidhub.com/8ZJdVFg7gz/2

[View certificate details](https://understandingwebpki.com/?cert=MIICoTCCAkegAwIBAgIQSFscNJOjGceOfFyuBLk%2BVzAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMTAeFw0yMjAzMzAxMjU0MTVaFw0yMzAzMzAxMjU0MTRaMGAxCzAJBgNVBAYTAlVTMRMwEQYDVQQKEwpUcmFuc05leHVzMRcwFQYDVQQLEw5GdXNpb24gQ29ubmVjdDEjMCEGA1UEAxMaRnVzaW9uIENvbm5lY3QgU0hBS0VOIDI3MjAwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAT63jwuwao3XgseWltjK2ryDbERqdg1YPJq5ERkdas77TTVUuYvquN%2FKmHQd5tGNxd0HFT74V7uRnF%2B0%2Faql%2B59o4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBQgC9%2Fhoul3TB6S5BONMObenIn0jjAfBgNVHSMEGDAWgBSUhjk%2F5PWSoJ%2F%2F3Cd1GppG8HnhYjAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMBYGCCsGAQUFBwEaBAowCKAGFgQyNzIwMAoGCCqGSM49BAMCA0gAMEUCIDNwXtIcpzboL8SV5qUwPjA3OiDBDuLVOeuxTqtMfwpQAiEAt7LcSEYFhKVzHAXNrJ9oJPpw8%2Fzrflp1JEJ6FudrYKU%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | the Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: [2.16.840.1.114569.1.1.3 2.16.840.1.114569.1.1.4] |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 21 Nov 23 17:53 UTC