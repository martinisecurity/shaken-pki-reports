# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Hawaiian Telcom SHAKEN Cert 009G

Tested At: 16 Mar 23 19:11 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 466 day(s)\
Subject: CN=Hawaiian Telcom SHAKEN Cert 009G, O=Hawaiian Telcom, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://sti-cr.cgah.tnsi.com/certs/3e5c8831f88b085823dc347977abc7c3b3c4ae74

[View certificate details](https://understandingwebpki.com/?cert=MIICWDCCAf%2BgAwIBAgIQOQUqOZEeJjGXLniu6d23pTAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDYyNDIyMzMzN1oXDTI0MDYyMzIyMzMzN1owUjELMAkGA1UEBhMCVVMxGDAWBgNVBAoMD0hhd2FpaWFuIFRlbGNvbTEpMCcGA1UEAwwgSGF3YWlpYW4gVGVsY29tIFNIQUtFTiBDZXJ0IDAwOUcwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARcJg18PBO4J1HRk8b8eGIxwzSmgW3P3K0v5EUBtwGzInfo5br1N6GzxMzBADLR49cDHXhIq9qTrVab%2FZ7v4dCZo4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgXgMBYGCCsGAQUFBwEaBAowCKAGFgQwMDlHMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFPUe76ZUTXN8P8fXUWX5QnJmpNecMB8GA1UdIwQYMBaAFM0epwAQENoyHWkaOdXSRgssPIfWMAoGCCqGSM49BAMCA0cAMEQCIGI0KLpElJY3QpGIh8pEBmkDbdQp4SEtACyyHiLv1GjpAiB%2BbuEUXQfWBUZbGLFlq6j80fkNR7e%2FYzGr9pdDg3J76A%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_key_usage](../../ISSUES/e_atis_key_usage/README.md) | error | ATIS1000080 | The Key Usage extension shall contain a single key usage value of digitalSignature |
| [e_atis_issuer_dn](../../ISSUES/e_atis_issuer_dn/README.md) | error | ATIS1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |

### Not Effective

- e_atis_extension_unknown
- e_atis_serial_number
- e_atis_signature_algorithm
- e_atis_subject_cn

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 16 Mar 23 19:18 UTC