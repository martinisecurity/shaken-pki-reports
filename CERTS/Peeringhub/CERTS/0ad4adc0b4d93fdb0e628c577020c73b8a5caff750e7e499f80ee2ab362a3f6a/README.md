# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Peeringhub Inc Root CA

Tested At: 26 Aug 24 18:03 UTC\
Initial Validity Period: 7300 day(s)\
Remaining Validity Period: 5952 day(s)\
Subject: CN=Peeringhub Inc Root CA, OU=Certification Authorities, O=Peeringhub Inc, C=US\
Issuer: CN=Peeringhub Inc Root CA, OU=Certification Authorities, O=Peeringhub Inc, C=US

[View certificate details](https://x509.io/?cert=MIICEDCCAbWgAwIBAgIJAJZM3gWl52yeMAoGCCqGSM49BAMCMGsxCzAJBgNVBAYTAlVTMRcwFQYDVQQKDA5QZWVyaW5naHViIEluYzEiMCAGA1UECwwZQ2VydGlmaWNhdGlvbiBBdXRob3JpdGllczEfMB0GA1UEAwwWUGVlcmluZ2h1YiBJbmMgUm9vdCBDQTAeFw0yMDEyMTcxNTMxMDRaFw00MDEyMTIxNTMxMDRaMGsxCzAJBgNVBAYTAlVTMRcwFQYDVQQKDA5QZWVyaW5naHViIEluYzEiMCAGA1UECwwZQ2VydGlmaWNhdGlvbiBBdXRob3JpdGllczEfMB0GA1UEAwwWUGVlcmluZ2h1YiBJbmMgUm9vdCBDQTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABOnrCNZEAlQT3yv%2FZ2pcR%2FNtTuNf7k3pjojYhYvA83uMCzdu9KlG%2BLdDwFcioZPF027ks4CaeenplBMGhNMkVnajQjBAMA4GA1UdDwEB%2FwQEAwIBhjAPBgNVHRMBAf8EBTADAQH%2FMB0GA1UdDgQWBBQc6aJG%2F%2BsTGtSS54L0Yo7%2BBaJ%2FzzAKBggqhkjOPQQDAgNJADBGAiEA%2Frriu3nwg8fan3oUiIA%2B12PM%2Bzuj6XfE7Ay92wNItTwCIQCrkOhIJG0L2Mg%2FwpDJPTWO4gaTpGqYbdexFlsinEAaFg%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_ca](../../ISSUES/e_atis_subject_cn_ca/README.md) | error | ATIS1000080 | Common Name attribute 'Peeringhub Inc Root CA' does not contain 'SHAKEN' |

### Not Effective

- e_atis_ext_certificate_policies_root
- e_atis_ext_crl_distribution_root
- e_atis_ext_not_specified_ca
- e_atis_serial_number_size_ca
- e_atis_subject_c_iso_ca
- e_atis_subject_cn_root
- e_atis_subject_key_identifier_size_ca
- e_atis_subject_o_required_ca

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 26 Aug 24 18:03 UTC