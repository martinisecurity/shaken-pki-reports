# STIR/SHAKEN CA Ecosystem Compliance

[Approved Certificate Authorities](https://authenticate.iconectiv.com/approved-certification-authorities) in the STIR/SHAKEN ecosystem are required to meet technical requirements from [ATIS-1000080](https://access.atis.org/apps/group_public/document.php?document_id=62163) and policy requirements from the supporting CA ecosystem’s [Certificate Policy](https://authenticate.iconectiv.com/documents-authenticate).

This report is generated using [Zlint](https://github.com/zmap/zlint) a tool commonly used to asses CA ecosystem compliance with such requirements. The tests used to generate this report are currently not part of the main Zlint distribution but can be found [here](https://github.com/martinisecurity/zlint).

## Summary


### CA Certificates


| Issuers | Certificates | Errors | Warnings | Notices | Not Effective |
|---------|--------------|--------|----------|---------|---------------|
| [Comcast](Comcast/README.md#ca-certificates) | 1 (7.69%) | 0 (0.00%) | 0 (0.00%) | 1 (100.00%) | 1 (100.00%) |
| [GBSDTech](GBSDTech/README.md#ca-certificates) | 1 (7.69%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) | 1 (100.00%) |
| [Martini Security](Martini_Security/README.md#ca-certificates) | 1 (7.69%) | 0 (0.00%) | 1 (100.00%) | 0 (0.00%) | 0 (0.00%) |
| [Metaswitch](Metaswitch/README.md#ca-certificates) | 1 (7.69%) | 1 (100.00%) | 0 (0.00%) | 0 (0.00%) | 1 (100.00%) |
| [NetNumber](NetNumber/README.md#ca-certificates) | 2 (15.38%) | 1 (50.00%) | 1 (50.00%) | 0 (0.00%) | 2 (100.00%) |
| [Neustar](Neustar/README.md#ca-certificates) | 2 (15.38%) | 0 (0.00%) | 0 (0.00%) | 2 (100.00%) | 2 (100.00%) |
| [Peeringhub](Peeringhub/README.md#ca-certificates) | 1 (7.69%) | 0 (0.00%) | 0 (0.00%) | 1 (100.00%) | 1 (100.00%) |
| [Ribbon Communications](Ribbon_Communications/README.md#ca-certificates) | 1 (7.69%) | 0 (0.00%) | 0 (0.00%) | 1 (100.00%) | 1 (100.00%) |
| [Sansay](Sansay/README.md#ca-certificates) | 1 (7.69%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) | 1 (100.00%) |
| [T-Mobile](T-Mobile/README.md#ca-certificates) | 1 (7.69%) | 0 (0.00%) | 0 (0.00%) | 1 (100.00%) | 1 (100.00%) |
| [TransNexus](TransNexus/README.md#ca-certificates) | 1 (7.69%) | 1 (100.00%) | 1 (100.00%) | 1 (100.00%) | 1 (100.00%) |
| **Total** | 13 (100%) | 3 (23.08%) | 3 (23.08%) | 7 (53.85%) | 12 (92.31%) |

\* The percent of certificates per issuer is calculated against total certificates from all issuers.\
\*\* The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer.\
\*\*\* 0 certificates skipped because they are currently expired.\
\*\*\*\* 0 certificates skipped because they are not currently trusted by the STI-PA.

## Key

| Type | Description |
|------|-------------|
| Error | Tests in which the specifications are unambiguous on what the expected behavior must be. |
| Warning	| Tests in which the specifications are ambiguous or are provide only a recommendation. |
| Notice | Tests in which industry best practices are not followed. |
| Not Effective	| Tests that exist in the current specifications but were not in effect at the time of issuance. |

Generated: 27/10/2022 at 22:13:25