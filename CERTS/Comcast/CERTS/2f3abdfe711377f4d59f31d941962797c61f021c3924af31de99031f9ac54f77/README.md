# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Comcast SHAKEN Intermediate CA

Tested At: 01 Nov 22 20:34 UTC\
Initial Validity Period: 7000 day(s)\
Remaining Validity Period: 6061 day(s)\
Subject: CN=Comcast SHAKEN Intermediate CA, O=Comcast, ST=Pennsylvania, C=US\
Issuer: CN=Comcast SHAKEN Root CA, O=Comcast, L=Philadelphia, ST=Pennsylvania, C=US

View: [Click to view](https://understandingwebpki.com/?cert=MIICIDCCAcagAwIBAgICEAIwCgYIKoZIzj0EAwIwbjELMAkGA1UEBhMCVVMxFTATBgNVBAgMDFBlbm5zeWx2YW5pYTEVMBMGA1UEBwwMUGhpbGFkZWxwaGlhMRAwDgYDVQQKDAdDb21jYXN0MR8wHQYDVQQDDBZDb21jYXN0IFNIQUtFTiBSb290IENBMB4XDTIwMDQwNjEzNDgzMFoXDTM5MDYwNjEzNDgzMFowXzELMAkGA1UEBhMCVVMxFTATBgNVBAgMDFBlbm5zeWx2YW5pYTEQMA4GA1UECgwHQ29tY2FzdDEnMCUGA1UEAwweQ29tY2FzdCBTSEFLRU4gSW50ZXJtZWRpYXRlIENBMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEFm8KIZEQBTAyvtTSQGrhSLUteUroEc7BOk8jjsnRK68HJiJRnNF%2BymUObx8t0zQ89YTrIu8Kdn91JEIBBtfwK6NjMGEwHQYDVR0OBBYEFL%2BOHGBVJMURSM9CRUS%2FsPQ6K%2BYoMB8GA1UdIwQYMBaAFJGQyrGGDk8WXr61N1E%2FaXnlIxscMA8GA1UdEwEB%2FwQFMAMBAf8wDgYDVR0PAQH%2FBAQDAgIEMAoGCCqGSM49BAMCA0gAMEUCIQCDqdriNDQQRzliTU%2FManJXY74z2NMQjmrRn4Uld180LQIgJDpxoJ3uHAH833qoy%2FacrVxSaegS9Em8Mc3h8AwXXuM%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_ca_subject_rdn_unknown](../../ISSUES/w_shaken_ca_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, L, and O should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


### Not Effective

- e_atis_basic_constraints
- e_atis_ca_authority_key_identifier
- e_atis_ca_certificate_policies
- e_atis_ca_crl_distribution
- e_atis_ca_extension_unknown
- e_atis_ca_issuer_dn
- e_atis_ca_key_usage
- e_atis_ca_serial_number
- e_atis_ca_signature_algorithm
- e_atis_ca_subject
- e_atis_ca_subject_cn
- e_atis_ca_subject_key_identifier
- e_atis_ca_subject_public_key
- e_atis_ca_version
- e_us_cp_ca_key_usage_crl_sign

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 01/11/2022 at 20:34:21