using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using TMPro;
using UnityEngine.UI;

public class PuzzleController : MonoBehaviour
{
    public TMP_InputField input;
    public Animator door;
    public ParticleSystem p;
    
    void Start ()
    {
        //  var input = gameObject.GetComponent<InputField>();
        //  var se= new InputField.SubmitEvent();
        //  se.AddListener(SubmitAnswer);
        //  input.onEndEdit = se;

        //or simply use the line below, 
        // input.onEndEdit.AddListener(SubmitAnswer);  // This also works
    }
    
    public void SubmitAnswer(string arg0)
    {
        Debug.Log(arg0);
        if (string.Equals(arg0,"hack"))
        {
            p.Play();
            door.Play("Door Opened");
            GameManager.gm.LoadWithDelay("Main Menu",3);
        }
    }
}
