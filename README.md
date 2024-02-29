## Informazioni generali
Progetto di WASA (Web and Software Architecture) A.A.2023/24.
Questo progetto consiste nello sviluppo di un sito web con l'uso delle seguenti tecnologie: API (OpenAPI), Backend (GoLang, SQLite), Frontend (Vue js, Javascript) e Docker.
Nella fattispecie il sito è una sorta di social (facsimile Facebook, Instagram, etc.), dove ogni utente mediante l'accesso o la registrazione con un nickname, può pubblicare post (foto), seguire altri utenti presenti e interagire con i loro contenuti mediante like e commenti. 

## Funzionalità
***
Nello specifico, le funzioni implentate sono le seguenti:
 * doLogin 
 * setMyUserName
 * uploadPhoto
 * followUser
 * unfollowUser
 * banUser
 *  unbanUser
 *  etUserProfile
 * getMyStream
 * likePhoto
 * unlikePhoto
 * commentPhoto
 * uncommentPhoto
 * deletePhoto
***

## Come avviare Backend e Frontend (Linux & WSL)
***
* ATTENZIONE: Consiglio vivamente di utilizzare dei sistemi Linux oppure di utilizzare WSL (su Win10/Win11), in quanto i comandi che seguiranno sono specifici di Linux, almeno in parte.
* CONSIGLIO: Se volete usare Windows 10/11, scaricate WSL, che vi permette di avere un sottosistema Ubuntu. Questo è molto comodo per evitare di installare un altro sistema operativo, inoltre, è 
  molto semplice da usare anche VSCode in quanto si può collegare direttamente a WSL.
* Per prima cosa, andiamo a scaricare la repo di github con gitclone (ATTENZIONE: non scaricate la repo come .zip altrimenti potrebbero esserci problemi con i permessi e potrebbe non funzionare)
* Installare npm (ci serverirà nella parte di frontend) e aggiornare
* Installare GoLang
* Installare NodeJS
* Vue JS
* Vite js
```
$ git clone git@github.com:Reiver56/Wasa.git
$ npm install
$ npm install npm@latest -g
$ sudo wget https://golang.org/dl/go1.16.linux-amd64.tar.gz
$ sudo apt install Node.js
$ npm install -g @vue/cli
$ npm init @vitejs/app
```
* Una volta scaricate, ci spostiamo nella directory del progetto.
***
## Avviare le varie componenti:
***
* Avviare Backend:
  ```
  $ go run ./cmd/webapi
  ```
* Fatto questo comando, vi si aprirà la porta relativa al backend e dalla bash potrete anche testare le richieste HTTP

* Avviare Frontend:
  ```
  $ ./open-npm.sh
  ```
  * Eseguendo questo comando andremo ad aprire il container di npm e possiamo fare due cose:
  	* Avviare in dev mode:
      ```
      npm run dev
      ```
    * Fare un build della production e avviare (consigliata per verificare eventuali bug che non si vedono in dev mode):
      ```
      npm run build-prod
      npm run preview
      ```
  * Una volta eseguito questo il frontend sarà avviato e potrete utilizzare il sito e vederne le funzionalità implementate.
***
  
