export interface IResponse {
    total_count: number;
    incomplete_results: boolean;
    items: IItem[];
}

export interface IItem {
    Name: string;
    Description: string;
    html_url: string;
    stargazers_count: number;
    owner : Owner;
}

export class Owner {
    login: string;
}