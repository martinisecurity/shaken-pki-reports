# STIR/SHAKEN CA Ecosystem Compliance
## Peeringhub

### Certificate d2590d58fcb7193f05fdefc76a1897333a060627
Tested At: 2022-10-28 18:15:09 +0000 UTC\
Initial Validity Period: 7300 day(s)\
Remaining Validity Period: 6620 day(s)\
Subject: CN=Peeringhub Inc Root CA, OU=Certification Authorities, O=Peeringhub Inc, C=US\
Issuer: CN=Peeringhub Inc Root CA, OU=Certification Authorities, O=Peeringhub Inc, C=US

Link: 

View: [Click to view](https://understandingwebpki.com/?cert=MIICEDCCAbWgAwIBAgIJAJZM3gWl52yeMAoGCCqGSM49BAMCMGsxCzAJBgNVBAYTAlVTMRcwFQYDVQQKDA5QZWVyaW5naHViIEluYzEiMCAGA1UECwwZQ2VydGlmaWNhdGlvbiBBdXRob3JpdGllczEfMB0GA1UEAwwWUGVlcmluZ2h1YiBJbmMgUm9vdCBDQTAeFw0yMDEyMTcxNTMxMDRaFw00MDEyMTIxNTMxMDRaMGsxCzAJBgNVBAYTAlVTMRcwFQYDVQQKDA5QZWVyaW5naHViIEluYzEiMCAGA1UECwwZQ2VydGlmaWNhdGlvbiBBdXRob3JpdGllczEfMB0GA1UEAwwWUGVlcmluZ2h1YiBJbmMgUm9vdCBDQTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABOnrCNZEAlQT3yv%2FZ2pcR%2FNtTuNf7k3pjojYhYvA83uMCzdu9KlG%2BLdDwFcioZPF027ks4CaeenplBMGhNMkVnajQjBAMA4GA1UdDwEB%2FwQEAwIBhjAPBgNVHRMBAf8EBTADAQH%2FMB0GA1UdDgQWBBQc6aJG%2F%2BsTGtSS54L0Yo7%2BBaJ%2FzzAKBggqhkjOPQQDAgNJADBGAiEA%2Frriu3nwg8fan3oUiIA%2B12PM%2Bzuj6XfE7Ay92wNItTwCIQCrkOhIJG0L2Mg%2FwpDJPTWO4gaTpGqYbdexFlsinEAaFg%3D%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| n_pki_ca_key_usage | notice | SHAKEN PKI Best Practice | For CA certificates, the Key Usage extension should contain a single key usage value of keyCertSign |
| w_pki_ca_subject_rdn_unknown | warn | SHAKEN PKI Best Practice | Only CN, C, L, and O should be included. Additional RNDs may introduce ambiguity and may not be verifiable |

### Not Effective

- e_sti_ca_serial_number
- e_sti_root_extension_unknown
- e_sti_root_certificate_policies
- e_sti_ca_subject_cn
- e_cp1_3_ca_key_usage_crl_sign

\* Tests use the ATIS 1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.

Generated: 28/10/2022 at 18:15:47