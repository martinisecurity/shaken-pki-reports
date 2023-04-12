# STIR/SHAKEN CA Ecosystem Compliance

## Certificate American Broadband SHAKEN Cert 355D

Tested At: 12 Apr 23 01:45 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 389 day(s)\
Subject: CN=American Broadband SHAKEN Cert 355D, O=American Broadband, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://sti-cr.cgah.tnsi.com/certs/f786e53509092c8a45b19b7dcd6886b5316d333b

[View certificate details](https://understandingwebpki.com/?cert=MIICXzCCAgWgAwIBAgIQXb5l%2FXT4Re1URPM7XE5hSDAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDUwNTIxMDI1NFoXDTI0MDUwNDIxMDI1NFowWDELMAkGA1UEBhMCVVMxGzAZBgNVBAoMEkFtZXJpY2FuIEJyb2FkYmFuZDEsMCoGA1UEAwwjQW1lcmljYW4gQnJvYWRiYW5kIFNIQUtFTiBDZXJ0IDM1NUQwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASHDsAID3%2Bu%2F7Xv5AU9r6PJ8FQi5rayLQAgj7zJko0jKnzdkrDpq1fjE%2BluclcipIXP3Pq%2F16jR7OoNmGDSwFWVo4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgXgMBYGCCsGAQUFBwEaBAowCKAGFgQzNTVEMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFL5xqjVymgE1zuEO5Jjv1%2FmSM50lMB8GA1UdIwQYMBaAFM0epwAQENoyHWkaOdXSRgssPIfWMAoGCCqGSM49BAMCA0gAMEUCIFtCzQhaLeaRYSGMvb59ioNELqw73bkpdk7DMVy5ttZ6AiEAkgfVA%2B4Htu8VNw2MAPmMkor9hXIIhlfrsHInobGlsQA%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_key_usage](../../ISSUES/e_atis_key_usage/README.md) | error | ATIS1000080 | The Key Usage extension shall contain a single key usage value of digitalSignature |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_issuer_dn](../../ISSUES/e_atis_issuer_dn/README.md) | error | ATIS1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |

### Not Effective

- e_atis_extension_unknown
- e_atis_serial_number
- e_atis_signature_algorithm
- e_atis_subject_cn

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 12 Apr 23 01:46 UTC