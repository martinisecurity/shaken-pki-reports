# STIR/SHAKEN CA Ecosystem Compliance
## Metaswitch

### Certificate 7deaf409e0c7859925f43a64b96cdd8b0eb1be89
Tested At: 2022-10-27 22:32:14 +0000 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 527 day(s)\
Subject: CN=U. S. Telepacific Corp SHAKEN 7453, O=U. S. Telepacific Corp, L=Los Angeles, ST=California, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1

Link: https://qcall.meta.tpx.net/certs/shaken_cacert.crt

View: [Click to view](https://understandingwebpki.com/?cert=MIICjjCCAjSgAwIBAgIQOs%2B8mLvpJuimIwt%2F8leOJzAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDQwNzE2NDQzNFoXDTI0MDQwNjE2NDQzNFowgYYxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRQwEgYDVQQHDAtMb3MgQW5nZWxlczEfMB0GA1UECgwWVS4gUy4gVGVsZXBhY2lmaWMgQ29ycDErMCkGA1UEAwwiVS4gUy4gVGVsZXBhY2lmaWMgQ29ycCBTSEFLRU4gNzQ1MzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABDemSJvLPHNy%2B4YLKb421Nhn4rZoqgNs05O3XtfuIXjCKM0TJ2Mmc8eXEttFFoGY7LNDSOe4hyp40tYh7oGU2uOjgdswgdgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCBeAwFgYIKwYBBQUHARoECjAIoAYWBDc0NTMwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUcxcPnhodcObxHUTHQVizJ4jhXSAwHwYDVR0jBBgwFoAUzR6nABAQ2jIdaRo51dJGCyw8h9YwCgYIKoZIzj0EAwIDSAAwRQIhAJnIqBIwlsBvII7ruDXkZCEMijzuFLQiYbfuNr%2F6pYz%2BAiAaWOSCo2ZO3na4cxjFryoc7WfK0gQuCCBW78xzMRW%2FjA%3D%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_sti_issuer_dn | error | ATIS-1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |
| e_sti_key_usage | error | ATIS-1000080 | The Key Usage extension shall contain a single key usage value of digitalSignature |

### Not Effective

- e_cp1_3_ambiguous_identifier
- e_sti_serial_number
- w_cp_1_3_subject_email
- e_sti_extension_unknown
- e_sti_signature_algorithm
- w_cp1_3_subject_rdn_unknown
- e_sti_subject_cn
- e_cp1_3_subject_sn

\* Tests use the ATIS 1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.

Generated: 27/10/2022 at 22:33:03