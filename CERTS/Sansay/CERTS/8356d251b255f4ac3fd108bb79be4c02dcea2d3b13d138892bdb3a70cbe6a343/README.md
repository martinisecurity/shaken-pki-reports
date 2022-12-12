# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Sansay Root CA US

Tested At: 12 Dec 22 23:45 UTC\
Initial Validity Period: 7300 day(s)\
Remaining Validity Period: 6457 day(s)\
Subject: CN=SHAKEN Sansay Root CA US, OU=Sansay CA, O=Sansay Corporation, L=San Diego, ST=California, C=US\
Issuer: CN=SHAKEN Sansay Root CA US, OU=Sansay CA, O=Sansay Corporation, L=San Diego, ST=California, C=US

[View certificate details](https://understandingwebpki.com/?cert=MIICUDCCAfWgAwIBAgIJAJkgyGgxR4kNMAoGCCqGSM49BAMCMIGKMQswCQYDVQQGEwJVUzETMBEGA1UECAwKQ2FsaWZvcm5pYTESMBAGA1UEBwwJU2FuIERpZWdvMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEhMB8GA1UEAwwYU0hBS0VOIFNhbnNheSBSb290IENBIFVTMB4XDTIwMDgyMTAxMjIwNFoXDTQwMDgxNjAxMjIwNFowgYoxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRIwEAYDVQQHDAlTYW4gRGllZ28xGzAZBgNVBAoMElNhbnNheSBDb3Jwb3JhdGlvbjESMBAGA1UECwwJU2Fuc2F5IENBMSEwHwYDVQQDDBhTSEFLRU4gU2Fuc2F5IFJvb3QgQ0EgVVMwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASOkTqogGQs88fIZEkSWFOrBiQHhV10YenZr86BmkYb131nPE8AJVF3diojafQYMKKFdP8qA4hssBxW9yW9iG8Po0IwQDAdBgNVHQ4EFgQUCq7%2FlvCbQaO9332%2FbdpFqOgEG7kwDwYDVR0TAQH%2FBAUwAwEB%2FzAOBgNVHQ8BAf8EBAMCAgQwCgYIKoZIzj0EAwIDSQAwRgIhAPTTm0wwyVjLrrwl20kbOgNMNVjAwLtcv6CnxWhoT20uAiEAuc9rBh4u%2Bs9GBKU6R2PXFnQ7ViZsZyZ5uyrP8VY%2Fdjo%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_ca_subject_rdn_unknown](../../ISSUES/w_shaken_ca_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, L, and O should be included. Additional RNDs may introduce ambiguity and may not be verifiable |

### Not Effective

- e_atis_basic_constraints
- e_atis_ca_issuer_dn
- e_atis_ca_key_usage
- e_atis_ca_serial_number
- e_atis_ca_signature_algorithm
- e_atis_ca_subject
- e_atis_ca_subject_cn
- e_atis_ca_subject_key_identifier
- e_atis_ca_subject_public_key
- e_atis_ca_version
- e_atis_root_certificate_policies
- e_atis_root_extension_unknown

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 12 Dec 22 23:45 UTC