# STIR/SHAKEN CA Ecosystem Compliance
## T-Mobile

### Certificate 45f4213a0916f509c15c42441aa6811f71c047ba
Tested At: 2022-10-28 18:54:33 +0000 UTC\
Initial Validity Period: 1826 day(s)\
Remaining Validity Period: 698 day(s)\
Subject: CN=TMOBILE-PROD-SUB-STIRSHAKEN-EC, O=TMOBILE-USA, C=US\
Issuer: CN=TMOBILE-PROD-ROOT-STIRSHAKEN-EC, O=TMOBILE-USA, C=US

Link: https://t-mobile-sticr.fosrvt.com/88a8e33055e725475530660e5d6c40d6adbe37ab7ae0ecc64b50205629548ae9.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIICBTCCAaqgAwIBAgICLLgwDAYIKoZIzj0EAwIFADBNMQswCQYDVQQGEwJVUzEUMBIGA1UEChMLVE1PQklMRS1VU0ExKDAmBgNVBAMTH1RNT0JJTEUtUFJPRC1ST09ULVNUSVJTSEFLRU4tRUMwHhcNMTkwOTI3MTcxMjA3WhcNMjQwOTI1MTc0MjA3WjBMMQswCQYDVQQGEwJVUzEUMBIGA1UEChMLVE1PQklMRS1VU0ExJzAlBgNVBAMTHlRNT0JJTEUtUFJPRC1TVUItU1RJUlNIQUtFTi1FQzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABDF0LDgsml4bEyfRuTdip0bHf0VP0%2FdOviKfiqGeuaAgGl8%2F7o7ppXOPIStancGkBgKTtMYB2oPgmpTidGwkI06jeTB3MB8GA1UdIwQYMBaAFIPpbpuk463zox84NDCefqFLuwnXMB0GA1UdDgQWBBQC2X9YhlA%2FeaB9iJtcgWc9pXwKeDAOBgNVHQ8BAf8EBAMCAaYwEQYDVR0gBAowCDAGBgRVHSAAMBIGA1UdEwEB%2FwQIMAYBAf8CAQAwDAYIKoZIzj0EAwIFAANHADBEAiAzm9KhK69YggTd3TxQyv1LSpMllq5oohwro6hCbU2VZQIga74k8Zllk2XkbYyz%2BMBy0mIRhBcrUIdKkk%2BwF2O1%2FBM%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| n_pki_ca_key_usage | notice | SHAKEN PKI Best Practice | For CA certificates, the Key Usage extension should contain a single key usage value of keyCertSign |

152 tests were ran and no warning or error level issues were found

### Not Effective

- e_sti_basic_constraints
- e_sti_ca_subject_cn
- e_sti_ca_key_usage
- e_sti_ca_subject
- e_sti_ca_subject_public_key
- e_sti_ca_serial_number
- e_sti_ca_authority_key_identifier
- e_sti_ca_version
- n_sti_ca_certificate_policy_critical
- e_cp1_3_ca_key_usage_crl_sign
- e_sti_ca_certificate_policies
- e_sti_ca_issuer_dn
- e_sti_ca_signature_algorithm
- e_sti_ca_extension_unknown
- e_sti_ca_subject_key_identifier
- e_sti_ca_crl_distribution

\* Tests use the ATIS 1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.

Generated: 28/10/2022 at 18:55:01