# STIR/SHAKEN CA Ecosystem Compliance

## Certificate U. S. Telepacific Corp SHAKEN 7453

Tested At: 31 Oct 22 19:21 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 523 day(s)\
Subject: CN=U. S. Telepacific Corp SHAKEN 7453, O=U. S. Telepacific Corp, L=Los Angeles, ST=California, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://qcall.meta.tpx.net/certs/shaken_cacert.crt

View: [Click to view](https://understandingwebpki.com/?cert=MIICjjCCAjSgAwIBAgIQOs%2B8mLvpJuimIwt%2F8leOJzAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDQwNzE2NDQzNFoXDTI0MDQwNjE2NDQzNFowgYYxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRQwEgYDVQQHDAtMb3MgQW5nZWxlczEfMB0GA1UECgwWVS4gUy4gVGVsZXBhY2lmaWMgQ29ycDErMCkGA1UEAwwiVS4gUy4gVGVsZXBhY2lmaWMgQ29ycCBTSEFLRU4gNzQ1MzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABDemSJvLPHNy%2B4YLKb421Nhn4rZoqgNs05O3XtfuIXjCKM0TJ2Mmc8eXEttFFoGY7LNDSOe4hyp40tYh7oGU2uOjgdswgdgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCBeAwFgYIKwYBBQUHARoECjAIoAYWBDc0NTMwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUcxcPnhodcObxHUTHQVizJ4jhXSAwHwYDVR0jBBgwFoAUzR6nABAQ2jIdaRo51dJGCyw8h9YwCgYIKoZIzj0EAwIDSAAwRQIhAJnIqBIwlsBvII7ruDXkZCEMijzuFLQiYbfuNr%2F6pYz%2BAiAaWOSCo2ZO3na4cxjFryoc7WfK0gQuCCBW78xzMRW%2FjA%3D%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_sti_key_usage](../../ISSUES/e_sti_key_usage/README.md) | error | ATIS&#x2011;1000080 | The Key Usage extension shall contain a single key usage value of digitalSignature |
| [w_pki_subject_rdn_unknown](../../ISSUES/w_pki_subject_rdn_unknown/README.md) | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_sti_issuer_dn](../../ISSUES/e_sti_issuer_dn/README.md) | error | ATIS&#x2011;1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |


### Not Effective

- e_cp1_3_ambiguous_identifier
- e_cp1_3_subject_sn
- e_sti_extension_unknown
- e_sti_serial_number
- e_sti_signature_algorithm
- e_sti_subject_cn
- w_cp_1_3_subject_email

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 31/10/2022 at 19:21:49