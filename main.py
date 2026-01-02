from fastapi import FastAPI, HTTPException
import subprocess


app = FastAPI(title = "Go Concurrency example | AFCON 2025 Group pairings")


@app.get("/groupa")
def afcon_groupa():
    afcon_groupa_teams = {"group_name": "groupa", "teams" : ["Zambia", "Morocco", "Comoros", "Mali"]}
    return afcon_groupa_teams 

@app.get("/groupb")
def afcon_groupb():
    afcon_groupb_teams = {"group_name": "groupb", "teams" : ["Egypt", "South_Africa", "Angola", "Zimbabwe"]}
    return afcon_groupb_teams 

@app.get("/groupc")
def afcon_groupc():
    afcon_groupc_teams = {"group_name": "groupc", "teams" : ["Nigeria", "Tunisia", "Tanzania", "Uganda"]}
    return afcon_groupc_teams 

@app.get("/groupd")
def afcon_groupd():
    afcon_groupd_teams = {"group_name": "groupd", "teams" : ["Senegal", "DR_Congo", "Benin", "Botswana"]}
    return afcon_groupd_teams 

@app.get("/groupe")
def afcon_groupe():
    afcon_groupe_teams = {"group_name": "groupd", "teams" : ["Algeria", "Burkina_Faso", "Sudan", "Equitorial_Guinea"]}
    return afcon_groupe_teams     

@app.get("/groupf")
def afcon_groupf():
    afcon_groupf_teams = {"group_name": "groupd", "teams" : ["Ivory_Coast", "Cameroon", "Mozambique", "Gabon"]}
    return afcon_groupf_teams   

@app.get("/afcon_groups_list_desc")
def run_afcon_groups_list_desc():
    try:
        result = subprocess.run(
            ["go", "run", "/app/groupteams_desc.go"],
            capture_output=True,
            text=True,
            check=True
        )

        teams = result.stdout.strip().splitlines()

        return {
            "afcon_teams": teams
        }

    except subprocess.CalledProcessError as e:
        raise HTTPException(
            status_code=500,
            detail={
                "stdout": e.stdout,
                "stderr": e.stderr,
                "returncode": e.returncode,
            },
        )


@app.get("/afcon_groups_list_aesc")
def run_afcon_groups_list_desc():
    try:
        result = subprocess.run(
            ["go", "run", "/app/groupteams_aesc.go"],
            capture_output=True,
            text=True,
            check=True
        )

        teams = result.stdout.strip().splitlines()

        return {
            "afcon_teams": teams
        }

    except subprocess.CalledProcessError as e:
        raise HTTPException(
            status_code=500,
            detail={
                "stdout": e.stdout,
                "stderr": e.stderr,
                "returncode": e.returncode,
            },
        )