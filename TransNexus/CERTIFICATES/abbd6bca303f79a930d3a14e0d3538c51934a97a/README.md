# STIR/SHAKEN CA Ecosystem Compliance
## TransNexus

### Certificate abbd6bca303f79a930d3a14e0d3538c51934a97a
Tested At: 2022-10-28 18:15:34 +0000 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 5 day(s)\
Subject: CN=Charter Communications Inc SHAKEN 5606, OU=Charter Communications Inc, O=TransNexus, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA1, OU=SHAKEN, O=TransNexus\\, Inc., C=US

Link: https://shaken.spectrum.com/cf1b3d3d-7f2b-42fd-a161-ebe61cd6565a.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIICuTCCAl%2BgAwIBAgIQemsdZ09BQKHxhrTLiMRjfzAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMTAeFw0yMTExMDIxNjQwMDFaFw0yMjExMDIxNjQwMDBaMHgxCzAJBgNVBAYTAlVTMRMwEQYDVQQKEwpUcmFuc05leHVzMSMwIQYDVQQLExpDaGFydGVyIENvbW11bmljYXRpb25zIEluYzEvMC0GA1UEAxMmQ2hhcnRlciBDb21tdW5pY2F0aW9ucyBJbmMgU0hBS0VOIDU2MDYwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAThIjWBWsXOYZCQVZi6Z%2F5OORvweOg3zN7OK1YvNF9UgiB8smcLv15itEsyyytSmVtbhwXG%2Bdn1OkxCAqgmGqBao4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBStEfznfqjnVndoPv%2Bri8BxpT4yCTAfBgNVHSMEGDAWgBSUhjk%2F5PWSoJ%2F%2F3Cd1GppG8HnhYjAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMBYGCCsGAQUFBwEaBAowCKAGFgQ1NjA2MAoGCCqGSM49BAMCA0gAMEUCIB4J6hjloqUbJgt1IMUc8x4kSxBLEj%2B0ZXBY6Xh2VYzoAiEAwQeeF%2BmK33JB0IxUdqXRt9DSD%2FiO6eE2e6eRa6DDQUQ%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| w_pki_subject_rdn_unknown | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| e_incorrect_ku_encoding | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |

### Not Effective

- e_sti_serial_number
- w_cp_1_3_subject_email
- e_sti_extension_unknown
- e_sti_signature_algorithm
- e_sti_subject_cn
- e_cp1_3_subject_sn
- e_cp1_3_ambiguous_identifier

\* Tests use the ATIS 1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.

Generated: 28/10/2022 at 18:15:47