# STIR/SHAKEN CA Ecosystem Compliance

## Certificate U. S. Telepacific Corp SHAKEN 7453

Tested At: 02 Nov 22 15:33 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 522 day(s)\
Subject: CN=U. S. Telepacific Corp SHAKEN 7453, O=U. S. Telepacific Corp, L=Los Angeles, ST=California, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://qcall.meta.tpx.net/certs/shaken_cacert.crt

View: [View certificate details](https://understandingwebpki.com/?cert=MIICjjCCAjSgAwIBAgIQOs%2B8mLvpJuimIwt%2F8leOJzAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDQwNzE2NDQzNFoXDTI0MDQwNjE2NDQzNFowgYYxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRQwEgYDVQQHDAtMb3MgQW5nZWxlczEfMB0GA1UECgwWVS4gUy4gVGVsZXBhY2lmaWMgQ29ycDErMCkGA1UEAwwiVS4gUy4gVGVsZXBhY2lmaWMgQ29ycCBTSEFLRU4gNzQ1MzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABDemSJvLPHNy%2B4YLKb421Nhn4rZoqgNs05O3XtfuIXjCKM0TJ2Mmc8eXEttFFoGY7LNDSOe4hyp40tYh7oGU2uOjgdswgdgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCBeAwFgYIKwYBBQUHARoECjAIoAYWBDc0NTMwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUcxcPnhodcObxHUTHQVizJ4jhXSAwHwYDVR0jBBgwFoAUzR6nABAQ2jIdaRo51dJGCyw8h9YwCgYIKoZIzj0EAwIDSAAwRQIhAJnIqBIwlsBvII7ruDXkZCEMijzuFLQiYbfuNr%2F6pYz%2BAiAaWOSCo2ZO3na4cxjFryoc7WfK0gQuCCBW78xzMRW%2FjA%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_key_usage](../../ISSUES/e_atis_key_usage/README.md) | error | ATIS1000080 | The Key Usage extension shall contain a single key usage value of digitalSignature |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_issuer_dn](../../ISSUES/e_atis_issuer_dn/README.md) | error | ATIS1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |

### Not Effective

- e_atis_extension_unknown
- e_atis_serial_number
- e_atis_signature_algorithm
- e_atis_subject_cn

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 02 Nov 22 15:41 UTC