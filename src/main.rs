use serde_json::{Result, Value};
use std::fs::File;
use std::io;
use std::env;

fn rem_first_and_last(value: &str) -> &str {
    let mut chars = value.chars();
    chars.next();
    chars.next_back();
    chars.as_str()
}

fn main() -> Result<()>{
    let response = reqwest::blocking::get("https://raw.githubusercontent.com/pui4/compack/main/data.json").unwrap();
    let v: Value = serde_json::from_str(&response.text().unwrap())?;
    
    let args: Vec<String> = env::args().collect();
    if &args[1] == "install" {
        let resp = reqwest::blocking::get(rem_first_and_last(&v[&args[2]].to_string())).expect("request failed");
        let body = resp.text().expect("body invalid");
        let mut out = File::create("discord.AppImage").expect("failed to create file");
        io::copy(&mut body.as_bytes(), &mut out).expect("failed to copy content");
    }

    Ok(())
}