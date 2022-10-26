# STIR/SHAKEN CA Ecosystem Compliance
## TransNexus

### Certificate 7c82e3e6b326e4aef329019feebd5b6342b224ed
Tested At: 2022-10-26 21:00:11 +0000 UTC\
Initial Validity Period: 4383 day(s)\
Remaining Validity Period: 3662 day(s)\
Subject: CN=TransNexus Issuing CA G2, OU=Certification Authorities, O=TransNexus, C=US\
Issuer: CN=TransNexus Root CA G2, OU=Certification Authorities, O=TransNexus, C=US

Link: https://certificates.transnexus.com/706J/9a9ea71d-1c8c-424c-8109-df524a4633d8.crt

View: [Click to view](https://understandingwebpki.com/?cert=MIICSjCCAe%2BgAwIBAgIQb1EQH5xFyuPdlc4wTK2DPDAKBggqhkjOPQQDAjBmMQswCQYDVQQGEwJVUzETMBEGA1UEChMKVHJhbnNOZXh1czEiMCAGA1UECxMZQ2VydGlmaWNhdGlvbiBBdXRob3JpdGllczEeMBwGA1UEAxMVVHJhbnNOZXh1cyBSb290IENBIEcyMB4XDTIwMTEwNDAwMDAwMFoXDTMyMTEwMzIzNTk1OVowaTELMAkGA1UEBhMCVVMxEzARBgNVBAoTClRyYW5zTmV4dXMxIjAgBgNVBAsTGUNlcnRpZmljYXRpb24gQXV0aG9yaXRpZXMxITAfBgNVBAMTGFRyYW5zTmV4dXMgSXNzdWluZyBDQSBHMjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABNlfglH3GiHuzkfpNRLQFH5ptp77%2BxCnIIPczCQ%2B84Aik%2BnOeDdjFU64pehF7jJSS6kMGzvsFjbE5vZzMKMfOqqjfDB6MA4GA1UdDwEB%2FwQEAwIBBjAPBgNVHRMBAf8EBTADAQH%2FMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQU%2FiVuNntTwyoXcwhCwWTGeYunJqgwHwYDVR0jBBgwFoAUd51fsT76PVNl361YAUbxtG5iZrYwCgYIKoZIzj0EAwIDSQAwRgIhAJLx79ihUjgBks75UUKVJ6Uo4KkUtWCT896FP6VRmegWAiEA9uaIZGFGabQNAVf1o6rAtN%2Fs03a6p%2BPGC62%2B0n5rUro%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| n_pki_ca_key_usage | notice | SHAKEN PKI Best Practice | For CA certificates, the Key Usage extension should contain a single key usage value of keyCertSign |

150 tests were ran and no warning or error level issues were found

### Not Effective

- e_cp1_3_ca_key_usage_crl_sign
- w_cp1_3_ca_subject_rdn_unknown
- e_sti_ca_certificate_policies
- e_sti_ca_subject
- e_sti_basic_constraints
- e_sti_ca_subject_public_key
- e_sti_ca_authority_key_identifier
- e_sti_ca_issuer_dn
- e_sti_ca_version
- e_sti_ca_signature_algorithm
- e_sti_ca_key_usage
- n_sti_ca_certificate_policy_critical
- e_sti_ca_crl_distribution
- e_sti_ca_subject_key_identifier
- e_sti_ca_extension_unknown
- e_sti_ca_subject_cn
- e_sti_ca_serial_number

\* Tests use the ATIS 1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.

Generated: 26/10/2022 at 21:01:13