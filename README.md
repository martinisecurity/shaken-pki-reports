# STIR/SHAKEN CA Ecosystem Compliance

[Approved Certificate Authorities](https://ecosystemcompliance.martinisecurity.com/#:~:text=Approved%20Certificate%20Authorities) in the STIR/SHAKEN ecosystem are required to meet technical requirements from [ATIS-1000080](https://access.atis.org/apps/group_public/document.php?document_id=62163) and policy requirements from the supporting CA ecosystem’s [Certificate Policy](https://authenticate.iconectiv.com/documents-authenticate).

This report is broken int two parts:
1. One generated using [Zlint](https://github.com/zmap/zlint) a tool commonly used to asses CA ecosystem compliance with such requirements. The tests used to generate this report are currently not part of the main [Zlint](https://github.com/martinisecurity/zlint) distribution but can be found here.
2. One generated with a custom script that eumerates the known STIR/SHAKEN certificates and asses each repository against the current rule set . The source for this test can be found here while the report itself can be found [here](REPOS/README.md).

## Summary

### Leaf Certificates

- 1103 certificates were included in the corpus being tested
- 146 certificates in the corpus were skipped because they are duplicates
- 537 certificates in the corpus were skipped because they are expired
- 4 certificates in the corpus were skipped because they are not currently trusted
- 416 certificates being tested against the remaining rules
- 3.73 issues on average found in unexpired, trusted, and non-compliant certificates
- 92.31% of certificates contain one or more Error level issue
- 49.28% of certificates contain one or more Warning level issue
- 2.16% of certificates contain one or more Notice level issue
- 14.18% of certificates are too old to be assessed against currently enforced expectations
- 279 days is the average remaining validity for the certificates in the corpus
- 280 days is the average initial validity for the certificates in the corpus
- 199 certificates expire in the next 30 days
- 3.59 average number of unexpired certificates per OCN observed
- 307 unique OCNs observed in unexpired and valid certificate corpus

### CA Certificates

- 37 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 1 certificates in the corpus were skipped because they are expired
- 3 certificates in the corpus were skipped because they are not currently trusted
- 33 certificates being tested against the remaining rules
- 2.05 issues on average found in unexpired, trusted, and non-compliant certificates
- 51.52% of certificates contain one or more Error level issue
- 51.52% of certificates contain one or more Warning level issue
- 3.03% of certificates contain one or more Notice level issue
- 69.70% of certificates are too old to be assessed against currently enforced expectations
- 5681 days is the average remaining validity for the certificates in the corpus
- 5612 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

## Certificate Repository URL

- 45.67% of certificate repository URLs contain one or more Error level issue
- 71.15% of certificates repository URLs contain one or more Warning level issue
- 0.00% of certificates repository URLs contain one or more Notice level issue

## Details

\* The percent of certificates per issuer is calculated against total certificates from all issuers.\
\*\* The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer.\
\*\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.

### Leaf Certificates

| Issuers | Certificates | Errors | Warnings | Notices | Not Effective |
|---------|--------------|--------|----------|---------|---------------|
| [Comcast](CERTS/Comcast/README.md#leaf-certificates) | 31 (7.45%) | 31 (100.00%) | 31 (100.00%) | 0 (0.00%) | 0 (0.00%) |
| [GBSDTech](CERTS/GBSDTech/README.md#leaf-certificates) | 2 (0.48%) | 2 (100.00%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) |
| [Martini Security](CERTS/Martini_Security/README.md#leaf-certificates) | 28 (6.73%) | 0 (0.00%) | 0 (0.00%) | 5 (17.86%) | 0 (0.00%) |
| [Metaswitch](CERTS/Metaswitch/README.md#leaf-certificates) | 30 (7.21%) | 30 (100.00%) | 1 (3.33%) | 0 (0.00%) | 28 (93.33%) |
| [NetNumber](CERTS/NetNumber/README.md#leaf-certificates) | 4 (0.96%) | 4 (100.00%) | 0 (0.00%) | 4 (100.00%) | 0 (0.00%) |
| [Neustar](CERTS/Neustar/README.md#leaf-certificates) | 106 (25.48%) | 102 (96.23%) | 31 (29.25%) | 0 (0.00%) | 30 (28.30%) |
| [Peeringhub](CERTS/Peeringhub/README.md#leaf-certificates) | 8 (1.92%) | 8 (100.00%) | 1 (12.50%) | 0 (0.00%) | 0 (0.00%) |
| [Ribbon Communications](CERTS/Ribbon_Communications/README.md#leaf-certificates) | 3 (0.72%) | 3 (100.00%) | 3 (100.00%) | 0 (0.00%) | 0 (0.00%) |
| [Sansay](CERTS/Sansay/README.md#leaf-certificates) | 111 (26.68%) | 111 (100.00%) | 111 (100.00%) | 0 (0.00%) | 0 (0.00%) |
| [T-Mobile](CERTS/T-Mobile/README.md#leaf-certificates) | 2 (0.48%) | 2 (100.00%) | 2 (100.00%) | 0 (0.00%) | 0 (0.00%) |
| [TransNexus](CERTS/TransNexus/README.md#leaf-certificates) | 91 (21.88%) | 91 (100.00%) | 25 (27.47%) | 0 (0.00%) | 1 (1.10%) |
| **Total** | 416 (100.00%) | 384 (92.31%) | 205 (49.28%) | 9 (2.16%) | 59 (14.18%) |

### CA Certificates

| Issuers | Certificates | Errors | Warnings | Notices | Not Effective |
|---------|--------------|--------|----------|---------|---------------|
| [Comcast](CERTS/Comcast/README.md#ca-certificates) | 2 (6.06%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) | 2 (100.00%) |
| [GBSDTech](CERTS/GBSDTech/README.md#ca-certificates) | 2 (6.06%) | 1 (50.00%) | 0 (0.00%) | 0 (0.00%) | 2 (100.00%) |
| [Martini Security](CERTS/Martini_Security/README.md#ca-certificates) | 3 (9.09%) | 0 (0.00%) | 0 (0.00%) | 1 (33.33%) | 0 (0.00%) |
| [Metaswitch](CERTS/Metaswitch/README.md#ca-certificates) | 2 (6.06%) | 2 (100.00%) | 0 (0.00%) | 0 (0.00%) | 2 (100.00%) |
| [NetNumber](CERTS/NetNumber/README.md#ca-certificates) | 3 (9.09%) | 2 (66.67%) | 2 (66.67%) | 0 (0.00%) | 3 (100.00%) |
| [Neustar](CERTS/Neustar/README.md#ca-certificates) | 6 (18.18%) | 4 (66.67%) | 6 (100.00%) | 0 (0.00%) | 4 (66.67%) |
| [Peeringhub](CERTS/Peeringhub/README.md#ca-certificates) | 2 (6.06%) | 2 (100.00%) | 2 (100.00%) | 0 (0.00%) | 1 (50.00%) |
| [Ribbon Communications](CERTS/Ribbon_Communications/README.md#ca-certificates) | 2 (6.06%) | 2 (100.00%) | 2 (100.00%) | 0 (0.00%) | 2 (100.00%) |
| [Sansay](CERTS/Sansay/README.md#ca-certificates) | 2 (6.06%) | 0 (0.00%) | 2 (100.00%) | 0 (0.00%) | 1 (50.00%) |
| [T-Mobile](CERTS/T-Mobile/README.md#ca-certificates) | 4 (12.12%) | 1 (25.00%) | 0 (0.00%) | 0 (0.00%) | 3 (75.00%) |
| [TransNexus](CERTS/TransNexus/README.md#ca-certificates) | 5 (15.15%) | 3 (60.00%) | 3 (60.00%) | 0 (0.00%) | 3 (60.00%) |
| **Total** | 33 (100.00%) | 17 (51.52%) | 17 (51.52%) | 1 (3.03%) | 23 (69.70%) |

### Key

| Type | Description |
|------|-------------|
| Errors | Tests in which the specifications are unambiguous on what the expected behavior must be. |
| Warnings | Tests in which the specifications are ambiguous or are provide only a recommendation. |
| Notices | Tests in which industry best practices are not followed. |
| Not Effective | Tests that exist in the current specifications but were not in effect at the time of issuance. |


Generated: 23 Nov 22 18:09 UTC