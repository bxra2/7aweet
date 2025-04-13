import type { Source } from "./source"
import type { Domain } from "./domain"

export type Term = {
    ID: number
    English: string
    French: string
    German: string
    Arabic: string
    Description: string
    URL: string

    // Foreign Key Fields
    SourceID: Number
    DomainID: Number

    // Associations
    Source: Source
    Domain: Domain
}