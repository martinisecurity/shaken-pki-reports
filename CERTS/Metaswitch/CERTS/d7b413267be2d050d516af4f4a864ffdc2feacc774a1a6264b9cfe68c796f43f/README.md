# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Verizon SHAKEN cert 5807

Tested At: 31 Oct 22 16:42 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 474 day(s)\
Subject: CN=Verizon SHAKEN cert 5807, OU=NNO CDS, O=Verizon Data Services LLC, L=Southlake, ST=Texas, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://sti.verizon.com/vzwcert/vzshaken-02-2024.crt

View: [Click to view](https://understandingwebpki.com/?cert=MIICkzCCAjigAwIBAgIQG7m6vUazJ%2Fqvch0J1rR9MTAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDIxNjE3MzAxM1oXDTI0MDIxNjE3MzAxM1owgYoxCzAJBgNVBAYTAlVTMQ4wDAYDVQQIDAVUZXhhczESMBAGA1UEBwwJU291dGhsYWtlMSIwIAYDVQQKDBlWZXJpem9uIERhdGEgU2VydmljZXMgTExDMRAwDgYDVQQLDAdOTk8gQ0RTMSEwHwYDVQQDDBhWZXJpem9uIFNIQUtFTiBjZXJ0IDU4MDcwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASBDPoBKI2KeWsVaW9EPaArBrlbiS9%2BAhGAqCDTmkbt520L5IAtToF5rphsKp02yLvpq8Us8Ze9Zh4A%2FYCSZnAho4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgXgMBYGCCsGAQUFBwEaBAowCKAGFgQ1ODA3MEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFLzCDjxfuU3ObZVJBsf27qs8H5TlMB8GA1UdIwQYMBaAFM0epwAQENoyHWkaOdXSRgssPIfWMAoGCCqGSM49BAMCA0kAMEYCIQD1xziWpLrPzBbUyTDWfbafwn7F7YNQSSgcd3l6Yga2BgIhANpoInVtDBD1HC9biSLaQGiQyuBq9ROxyt0qdh9%2FIA7v)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_sti_issuer_dn](../../ISSUES/e_sti_issuer_dn/README.md) | error | ATIS&#x2011;1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |
| [w_pki_subject_rdn_unknown](../../ISSUES/w_pki_subject_rdn_unknown/README.md) | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_sti_key_usage](../../ISSUES/e_sti_key_usage/README.md) | error | ATIS&#x2011;1000080 | The Key Usage extension shall contain a single key usage value of digitalSignature |


### Not Effective

- e_cp1_3_ambiguous_identifier
- e_cp1_3_subject_sn
- e_sti_extension_unknown
- e_sti_serial_number
- e_sti_signature_algorithm
- e_sti_subject_cn
- w_cp_1_3_subject_email

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 31/10/2022 at 16:43:22