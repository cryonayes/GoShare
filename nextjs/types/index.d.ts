// Implement API types etc.

interface UserFileModel {
    filename: string,
    file_type: string,
    file_size: number,
    access_code: string,
    owner: string,
    creation_date: string,
    shared: boolean
}

interface APIUserData {
    name: string,
    lastname: string,
    files: UserFileModel[]
}

interface APIResponseFiles {
    success: bool,
    message: string,
    data: APIUserData
}

interface APIResponseShare {
    success: bool,
    message: string,
    data: {
        accesslink: string
    }
}

interface APIResponseUpload {
    success: bool,
    message: string,
    data: UserFileModel
}

interface APIResponse {
    success: bool,
    message: string,
    data: any
}