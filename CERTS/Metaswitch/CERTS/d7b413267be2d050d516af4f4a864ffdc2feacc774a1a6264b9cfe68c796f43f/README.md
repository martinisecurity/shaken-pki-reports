# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Verizon SHAKEN cert 5807

Tested At: 17 Nov 22 19:19 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 456 day(s)\
Subject: CN=Verizon SHAKEN cert 5807, OU=NNO CDS, O=Verizon Data Services LLC, L=Southlake, ST=Texas, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://sti.verizon.com/vzwcert/vzshaken-02-2024.crt

[View certificate details](https://understandingwebpki.com/?cert=MIICkzCCAjigAwIBAgIQG7m6vUazJ%2Fqvch0J1rR9MTAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDIxNjE3MzAxM1oXDTI0MDIxNjE3MzAxM1owgYoxCzAJBgNVBAYTAlVTMQ4wDAYDVQQIDAVUZXhhczESMBAGA1UEBwwJU291dGhsYWtlMSIwIAYDVQQKDBlWZXJpem9uIERhdGEgU2VydmljZXMgTExDMRAwDgYDVQQLDAdOTk8gQ0RTMSEwHwYDVQQDDBhWZXJpem9uIFNIQUtFTiBjZXJ0IDU4MDcwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASBDPoBKI2KeWsVaW9EPaArBrlbiS9%2BAhGAqCDTmkbt520L5IAtToF5rphsKp02yLvpq8Us8Ze9Zh4A%2FYCSZnAho4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgXgMBYGCCsGAQUFBwEaBAowCKAGFgQ1ODA3MEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFLzCDjxfuU3ObZVJBsf27qs8H5TlMB8GA1UdIwQYMBaAFM0epwAQENoyHWkaOdXSRgssPIfWMAoGCCqGSM49BAMCA0kAMEYCIQD1xziWpLrPzBbUyTDWfbafwn7F7YNQSSgcd3l6Yga2BgIhANpoInVtDBD1HC9biSLaQGiQyuBq9ROxyt0qdh9%2FIA7v)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_atis_key_usage](../../ISSUES/e_atis_key_usage/README.md) | error | ATIS1000080 | The Key Usage extension shall contain a single key usage value of digitalSignature |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_issuer_dn](../../ISSUES/e_atis_issuer_dn/README.md) | error | ATIS1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |

### Not Effective

- e_atis_extension_unknown
- e_atis_serial_number
- e_atis_signature_algorithm
- e_atis_subject_cn

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 17 Nov 22 19:20 UTC