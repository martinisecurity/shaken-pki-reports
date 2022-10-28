# STIR/SHAKEN CA Ecosystem Compliance

[Approved Certificate Authorities](https://authenticate.iconectiv.com/approved-certification-authorities) in the STIR/SHAKEN ecosystem are required to meet technical requirements from [ATIS-1000080](https://access.atis.org/apps/group_public/document.php?document_id=62163) and policy requirements from the supporting CA ecosystem’s [Certificate Policy](https://authenticate.iconectiv.com/documents-authenticate).

This report is generated using [Zlint](https://github.com/zmap/zlint) a tool commonly used to asses CA ecosystem compliance with such requirements. The tests used to generate this report are currently not part of the main Zlint distribution but can be found [here](https://github.com/martinisecurity/zlint).

## Summary

### Leaf Certificates

- Average validity span as of leaf certificates 307 days
- Percentage of leaf certificates expiring in the next 30 days is 47.33%

| Issuers | Certificates | Errors | Warnings | Notices | Not Effective |
|---------|--------------|--------|----------|---------|---------------|
| [Comcast](Comcast/README.md#leaf-certificates) | 42 (12.32%) | 42 (100.00%) | 42 (100.00%) | 0 (0.00%) | 0 (0.00%) |
| [GBSDTech](GBSDTech/README.md#leaf-certificates) | 1 (0.29%) | 1 (100.00%) | 1 (100.00%) | 0 (0.00%) | 0 (0.00%) |
| [Martini Security](Martini_Security/README.md#leaf-certificates) | 7 (2.05%) | 0 (0.00%) | 0 (0.00%) | 4 (57.14%) | 0 (0.00%) |
| [Metaswitch](Metaswitch/README.md#leaf-certificates) | 21 (6.16%) | 21 (100.00%) | 6 (28.57%) | 0 (0.00%) | 20 (95.24%) |
| [NetNumber](NetNumber/README.md#leaf-certificates) | 4 (1.17%) | 4 (100.00%) | 1 (25.00%) | 4 (100.00%) | 0 (0.00%) |
| [Neustar](Neustar/README.md#leaf-certificates) | 97 (28.45%) | 82 (84.54%) | 33 (34.02%) | 0 (0.00%) | 33 (34.02%) |
| [Peeringhub](Peeringhub/README.md#leaf-certificates) | 5 (1.47%) | 5 (100.00%) | 5 (100.00%) | 0 (0.00%) | 0 (0.00%) |
| [Ribbon Communications](Ribbon_Communications/README.md#leaf-certificates) | 3 (0.88%) | 3 (100.00%) | 3 (100.00%) | 0 (0.00%) | 0 (0.00%) |
| [Sansay](Sansay/README.md#leaf-certificates) | 82 (24.05%) | 82 (100.00%) | 82 (100.00%) | 0 (0.00%) | 0 (0.00%) |
| [T-Mobile](T-Mobile/README.md#leaf-certificates) | 1 (0.29%) | 1 (100.00%) | 1 (100.00%) | 0 (0.00%) | 0 (0.00%) |
| [TransNexus](TransNexus/README.md#leaf-certificates) | 78 (22.87%) | 78 (100.00%) | 78 (100.00%) | 0 (0.00%) | 2 (2.56%) |
| **Total** | 341 (100%) | 319 (93.55%) | 252 (73.90%) | 8 (2.35%) | 55 (16.13%) |

\* The percent of certificates per issuer is calculated against total certificates from all issuers.\
\*\* The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer.\
\*\*\* Tests use the ATIS 1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.\
\*\*\*\* For information on failed certificate repository retrievals see this [report](URL/README.md).\
\*\*\*\*\* 273 certificates skipped because they are currently expired.\
\*\*\*\*\*\* 1 certificates skipped because they are not currently trusted by the STI-PA.

### CA Certificates


| Issuers | Certificates | Errors | Warnings | Notices | Not Effective |
|---------|--------------|--------|----------|---------|---------------|
| [Comcast](Comcast/README.md#ca-certificates) | 2 (5.88%) | 0 (0.00%) | 2 (100.00%) | 1 (50.00%) | 2 (100.00%) |
| [GBSDTech](GBSDTech/README.md#ca-certificates) | 3 (8.82%) | 1 (33.33%) | 3 (100.00%) | 0 (0.00%) | 3 (100.00%) |
| [Martini Security](Martini_Security/README.md#ca-certificates) | 3 (8.82%) | 0 (0.00%) | 3 (100.00%) | 1 (33.33%) | 0 (0.00%) |
| [Metaswitch](Metaswitch/README.md#ca-certificates) | 2 (5.88%) | 2 (100.00%) | 0 (0.00%) | 0 (0.00%) | 2 (100.00%) |
| [NetNumber](NetNumber/README.md#ca-certificates) | 3 (8.82%) | 2 (66.67%) | 3 (100.00%) | 0 (0.00%) | 3 (100.00%) |
| [Neustar](Neustar/README.md#ca-certificates) | 10 (29.41%) | 4 (40.00%) | 10 (100.00%) | 10 (100.00%) | 8 (80.00%) |
| [Peeringhub](Peeringhub/README.md#ca-certificates) | 2 (5.88%) | 1 (50.00%) | 2 (100.00%) | 2 (100.00%) | 1 (50.00%) |
| [Ribbon Communications](Ribbon_Communications/README.md#ca-certificates) | 2 (5.88%) | 0 (0.00%) | 2 (100.00%) | 2 (100.00%) | 2 (100.00%) |
| [Sansay](Sansay/README.md#ca-certificates) | 2 (5.88%) | 0 (0.00%) | 2 (100.00%) | 1 (50.00%) | 1 (50.00%) |
| [T-Mobile](T-Mobile/README.md#ca-certificates) | 2 (5.88%) | 0 (0.00%) | 0 (0.00%) | 2 (100.00%) | 2 (100.00%) |
| [TransNexus](TransNexus/README.md#ca-certificates) | 3 (8.82%) | 3 (100.00%) | 3 (100.00%) | 3 (100.00%) | 3 (100.00%) |
| **Total** | 34 (100%) | 13 (38.24%) | 30 (88.24%) | 22 (64.71%) | 27 (79.41%) |

\* The percent of certificates per issuer is calculated against total certificates from all issuers.\
\*\* The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer.\
\*\*\* 2 certificates skipped because they are currently expired.\
\*\*\*\* 1 certificates skipped because they are not currently trusted by the STI-PA.

## Key

| Type | Description |
|------|-------------|
| Error | Tests in which the specifications are unambiguous on what the expected behavior must be. |
| Warning	| Tests in which the specifications are ambiguous or are provide only a recommendation. |
| Notice | Tests in which industry best practices are not followed. |
| Not Effective	| Tests that exist in the current specifications but were not in effect at the time of issuance. |

Generated: 28/10/2022 at 19:22:10