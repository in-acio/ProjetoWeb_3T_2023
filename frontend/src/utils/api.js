import { useToast } from 'vue-toastification'

const toast = useToast()

const BASE_URL = "http://localhost:8080/"

export function USER(path, body){
    return { 
        url: BASE_URL + 'users' + path,
        options: {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(body)
        }
     }
};

export function GET_REQUEST(path, token){
    return { 
        url: BASE_URL + path,
        options: {
            method: 'GET',
            headers: {
                'Authorization': `Bearer ${token}`
            },
        }
     }
};

export function POST_REQUEST(path, method, token, body){
    return { 
        url: BASE_URL + path,
        options: {
            method: method,
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${token}`
            },
            body: JSON.stringify(body),
        },
     }
}

export async function makeRequest(path, token) {
    try {
        const data = GET_REQUEST(path, token);
        const req = await fetch(data.url, data.options);
        const json = await req.json();

        if (!req.ok) {
            toast.info("A sessão atual expirou");
            store.commit('logout');
            router.push('/');
        }
        
        return json;
    } catch (error) {
        toast.error(error.message);
        throw error;
    }
  }