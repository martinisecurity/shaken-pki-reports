# STIR/SHAKEN CA Ecosystem Compliance

## Certificate U. S. Telepacific Corp SHAKEN 7453

Tested At: 21 Nov 23 01:49 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 138 day(s)\
Subject: CN=U. S. Telepacific Corp SHAKEN 7453, O=U. S. Telepacific Corp, L=Los Angeles, ST=California, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://qcall.meta.tpx.net/certs/shaken_cacert.crt

[View certificate details](https://understandingwebpki.com/?cert=MIICjjCCAjSgAwIBAgIQOs%2B8mLvpJuimIwt%2F8leOJzAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDQwNzE2NDQzNFoXDTI0MDQwNjE2NDQzNFowgYYxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRQwEgYDVQQHDAtMb3MgQW5nZWxlczEfMB0GA1UECgwWVS4gUy4gVGVsZXBhY2lmaWMgQ29ycDErMCkGA1UEAwwiVS4gUy4gVGVsZXBhY2lmaWMgQ29ycCBTSEFLRU4gNzQ1MzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABDemSJvLPHNy%2B4YLKb421Nhn4rZoqgNs05O3XtfuIXjCKM0TJ2Mmc8eXEttFFoGY7LNDSOe4hyp40tYh7oGU2uOjgdswgdgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCBeAwFgYIKwYBBQUHARoECjAIoAYWBDc0NTMwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUcxcPnhodcObxHUTHQVizJ4jhXSAwHwYDVR0jBBgwFoAUzR6nABAQ2jIdaRo51dJGCyw8h9YwCgYIKoZIzj0EAwIDSAAwRQIhAJnIqBIwlsBvII7ruDXkZCEMijzuFLQiYbfuNr%2F6pYz%2BAiAaWOSCo2ZO3na4cxjFryoc7WfK0gQuCCBW78xzMRW%2FjA%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_ext_key_usage_ee](../../ISSUES/e_atis_ext_key_usage_ee/README.md) | error | ATIS1000080 | The Key Usage extension for STI end-entity certificates shall contain a single key usage value of digitalSignature (0). |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_cn_spc
- e_atis_subject_key_identifier_size
- e_atis_subject_o_required
- e_atis_tn_auth_list_spc_format
- e_shaken_certificate_policies_id

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 21 Nov 23 01:55 UTC