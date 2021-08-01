// Implement API types etc.

interface UserFileModel {
    filename: string,
    file_type: string,
    file_size: number,
    owner: string,
    creation_date: string
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