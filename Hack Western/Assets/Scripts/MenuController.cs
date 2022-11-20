using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class MenuController : MonoBehaviour
{
   public Animator menuAnimator;
   public void Play(){
      GameManager.gm.Load("Game");
   }

   public void LevelMenu(){
      menuAnimator.Play("Levels");
   }

   public void GameMenu(){
      menuAnimator.Play("Hosting Game");
   }
}
